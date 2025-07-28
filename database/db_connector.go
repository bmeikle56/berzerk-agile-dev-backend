package database

import (
	"database/sql"
	_ "github.com/lib/pq"
	"fmt"
	"os"
	"bzdev/models"
	"encoding/json"
)

func ConnectDB() (*sql.DB, error) {
	url :=  os.Getenv("DB_URL")

	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, fmt.Errorf("sql.Open error: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("db.Ping error: %w", err)
	}

	return db, nil
}

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

	// append the new ticket
	userData.Tickets = append(userData.Tickets, newTicket)

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

func DeleteTicketByTitle(db *sql.DB, username string, title string) error {
	// fetch user data from db
	userData, err := FetchUserData(db, username)
	if err != nil {
		return err
	}

	// filter out the ticket with the matching title
	found := false
	updatedTickets := make([]models.Ticket, 0)
	for _, ticket := range userData.Tickets {
		if ticket.Title == title {
			found = true
			continue // skip the ticket we're deleting
		}
		updatedTickets = append(updatedTickets, ticket)
	}

	if !found {
		return fmt.Errorf("ticket with title %q not found", title)
	}

	userData.Tickets = updatedTickets

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

func DeleteAllTickets(db *sql.DB, username string) error {
	// fetch user data from db
	userData, err := FetchUserData(db, username)
	if err != nil {
		return err
	}

	// check if there are any tickets to delete
	if len(userData.Tickets) == 0 {
		return fmt.Errorf("no tickets found for user %q", username)
	}

	// clear the tickets slice
	userData.Tickets = []models.Ticket{}

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


func UpdateTicketStatus(db *sql.DB, username string, title string, newStatus string) error {
	// fetch user data from db
	userData, err := FetchUserData(db, username)
	if err != nil {
		return err
	}

	// find and update the ticket by title
	found := false
	for i := range userData.Tickets {
		if userData.Tickets[i].Title == title {
			userData.Tickets[i].Status = newStatus
			found = true
			break
		}
	}

	if !found {
		return fmt.Errorf("ticket with title %q not found", title)
	}

	// marshal the updated data
	updatedBytes, err := json.Marshal(userData)
	if err != nil {
		return fmt.Errorf("marshal error: %w", err)
	}

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