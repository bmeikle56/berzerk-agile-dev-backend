package database

import (
	"database/sql"
	_ "github.com/lib/pq"
	"fmt"
	"bzdev/models"
	"encoding/json"
)

func CheckIfUserExists(db *sql.DB, username string) (bool, error) {
	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM bzdevusers WHERE username = $1)`

	err := db.QueryRow(query, username).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}

func FetchUserData(db *sql.DB, username string) (models.UserData, error) {
	query := `
		SELECT data
		FROM bzdevusers
		WHERE username = $1
	`

	var dataBytes []byte
	var userData models.UserData

	err := db.QueryRow(query, username).Scan(&dataBytes)
	if err != nil {
		return userData, fmt.Errorf("query error: %w", err)
	}

	if len(dataBytes) > 0 {
		if err := json.Unmarshal(dataBytes, &userData); err != nil {
			return userData, fmt.Errorf("failed to unmarshal JSON: %w", err)
		}
	}

	return userData, nil
}

func AssignTicketToUser(db *sql.DB, username string, newTicket models.Ticket) error {
	// fetch user data from db
	userData, err := FetchUserData(db, username)
	if err != nil {
		return err
	}

	// flag to check if repo was found
	repoFound := false

	// iterate through repos to find matching one
	for i, repo := range userData.Repos {
		if repo.Repo == newTicket.Repo {
			userData.Repos[i].Tickets = append(userData.Repos[i].Tickets, newTicket)
			repoFound = true
			break
		}
	}

	// if repo not found, create new repo entry
	if !repoFound {
		userData.Repos = append(userData.Repos, models.Repo{
			Repo:    newTicket.Repo,
			Tickets: []models.Ticket{newTicket},
		})
	}

	// marshal updated data
	updatedBytes, err := json.Marshal(userData)
	if err != nil {
		return fmt.Errorf("failed to marshal updated data: %w", err)
	}

	// write back to the database
	queryUpdate := `
		UPDATE bzdevusers
		SET data = $1
		WHERE username = $2
	`
	_, err = db.Exec(queryUpdate, updatedBytes, username)
	if err != nil {
		return fmt.Errorf("failed to update user data: %w", err)
	}

	return nil
}

func DeleteTicketByKey(db *sql.DB, username string, key string) error {
	// fetch user data from db
	userData, err := FetchUserData(db, username)
	if err != nil {
		return err
	}

	found := false

	// loop through repos
	for repoIndex, repo := range userData.Repos {
		newTickets := make([]models.Ticket, 0, len(repo.Tickets))
		for _, ticket := range repo.Tickets {
			if ticket.Key == key {
				found = true
				continue // skip ticket to delete
			}
			newTickets = append(newTickets, ticket)
		}
		userData.Repos[repoIndex].Tickets = newTickets
	}

	if !found {
		return fmt.Errorf("ticket with key %q not found", key)
	}

	// marshal updated data
	updatedBytes, err := json.Marshal(userData)
	if err != nil {
		return fmt.Errorf("failed to marshal updated data: %w", err)
	}

	// update DB
	queryUpdate := `
		UPDATE bzdevusers
		SET data = $1
		WHERE username = $2
	`
	_, err = db.Exec(queryUpdate, updatedBytes, username)
	if err != nil {
		return fmt.Errorf("failed to update user data: %w", err)
	}

	return nil
}

func KillRepoByName(db *sql.DB, username, repoName string) error {
	// fetch user data
	userData, err := FetchUserData(db, username)
	if err != nil {
		return err
	}

	// filter out the repo to kill
	newRepos := make([]models.Repo, 0, len(userData.Repos))
	found := false
	for _, repo := range userData.Repos {
		if repo.Repo == repoName {
			found = true
			continue // skip this repo
		}
		newRepos = append(newRepos, repo)
	}

	if !found {
		return fmt.Errorf("repo %q not found for user %q", repoName, username)
	}

	userData.Repos = newRepos

	// marshal updated data
	updatedBytes, err := json.Marshal(userData)
	if err != nil {
		return fmt.Errorf("failed to marshal updated user data: %w", err)
	}

	// update DB
	query := `
		UPDATE bzdevusers
		SET data = $1
		WHERE username = $2
	`
	_, err = db.Exec(query, updatedBytes, username)
	if err != nil {
		return fmt.Errorf("failed to update user data in DB: %w", err)
	}

	return nil
}


func DeleteAllTickets(db *sql.DB, username string) error {
	// fetch user data from db
	userData, err := FetchUserData(db, username)
	if err != nil {
		return err
	}

	// check if there are any tickets at all
	hasTickets := false
	for _, repo := range userData.Repos {
		if len(repo.Tickets) > 0 {
			hasTickets = true
			break
		}
	}

	if !hasTickets {
		return fmt.Errorf("no tickets found for user %q", username)
	}

	// clear all tickets for every repo
	for i := range userData.Repos {
		userData.Repos[i].Tickets = []models.Ticket{}
	}

	// marshal updated data
	updatedBytes, err := json.Marshal(userData)
	if err != nil {
		return fmt.Errorf("failed to marshal updated data: %w", err)
	}

	// update DB
	queryUpdate := `
		UPDATE bzdevusers
		SET data = $1
		WHERE username = $2
	`
	_, err = db.Exec(queryUpdate, updatedBytes, username)
	if err != nil {
		return fmt.Errorf("failed to update user data: %w", err)
	}

	return nil
}

func UpdateTicketStatusByRepo(db *sql.DB, username string, repoName string, key string, newStatus string) error {
	// fetch user data from db
	userData, err := FetchUserData(db, username)
	if err != nil {
		return err
	}

	found := false

	// search for the specific repo
	for repoIndex := range userData.Repos {
		if userData.Repos[repoIndex].Repo == repoName {
			repo := &userData.Repos[repoIndex]

			// ensure we don’t already have another active ticket
			if newStatus == "active" {
				for _, t := range repo.Tickets {
					if t.Status == "active" && t.Key != key {
						return fmt.Errorf("ticket %q is already active in repo %q", t.Key, repoName)
					}
				}
			}

			// update ticket status
			for ticketIndex := range repo.Tickets {
				if repo.Tickets[ticketIndex].Key == key {
					repo.Tickets[ticketIndex].Status = newStatus
					found = true
					break
				}
			}

			if !found {
				break
			}

			// reorder: active ticket first (if any)
			var active *models.Ticket
			var reordered []models.Ticket

			for _, t := range repo.Tickets {
				if t.Status == "active" {
					if active != nil {
						// safeguard: shouldn’t happen, but just in case
						return fmt.Errorf("multiple active tickets found in repo %q", repoName)
					}
					copyT := t
					active = &copyT
				}
			}

			// place active ticket first
			if active != nil {
				reordered = append(reordered, *active)
			}

			for _, t := range repo.Tickets {
				if t.Status != "active" {
					reordered = append(reordered, t)
				}
			}

			repo.Tickets = reordered
			break
		}
	}

	if !found {
		return fmt.Errorf("ticket with key %q in repo %q not found", key, repoName)
	}

	// marshal updated data
	updatedBytes, err := json.Marshal(userData)
	if err != nil {
		return fmt.Errorf("marshal error: %w", err)
	}

	// update DB
	queryUpdate := `
		UPDATE bzdevusers
		SET data = $1
		WHERE username = $2
	`
	_, err = db.Exec(queryUpdate, updatedBytes, username)
	if err != nil {
		return fmt.Errorf("update error: %w", err)
	}

	return nil
}

func InsertUser(db *sql.DB, username string, password string) error {
	query := `
		INSERT INTO bzdevusers (username, password, data)
		VALUES ($1, $2, $3)
	`
	initialData := `{}`
	_, err := db.Exec(query, username, password, initialData)
	if err != nil {
		return fmt.Errorf("InsertUser error: %w", err)
	}

	return nil
}

func FetchPasswordForUser(db *sql.DB, username string) (string, error) {
	query := `
		SELECT password
		FROM bzdevusers
		WHERE username = $1
	`

	var hashedPassword string
	err := db.QueryRow(query, username).Scan(&hashedPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", fmt.Errorf("user not found")
		}
		return "", fmt.Errorf("query error: %w", err)
	}

	return hashedPassword, nil
}

func FetchUser(db *sql.DB, username string) (*models.User, error) {
	query := `
		SELECT id, username, password
		FROM bzdevusers
		WHERE username = $1
	`

	row := db.QueryRow(query, username)

	var user models.User
	err := row.Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("query error: %w", err)
	}

	return &user, nil
}