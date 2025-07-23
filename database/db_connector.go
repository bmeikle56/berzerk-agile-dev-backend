package database

import (
	"database/sql"
	_ "github.com/lib/pq"
	"fmt"
	"os"
	"bzdev/models"
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
	query := `SELECT EXISTS(SELECT 1 FROM users WHERE username = $1)`

	err := db.QueryRow(query, username).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}

func InsertUser(db *sql.DB, username string, password string, diamonds int) error {
	query := `
		INSERT INTO users (username, password, diamonds)
		VALUES ($1, $2, $3)
	`
	_, err := db.Exec(query, username, password, diamonds)
	if err != nil {
		return fmt.Errorf("InsertUser error: %w", err)
	}

	return nil
}

func FetchPasswordForUser(db *sql.DB, username string) (string, error) {
	query := `
		SELECT password
		FROM users
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
		FROM users
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