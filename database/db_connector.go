package database

import (
	"database/sql"
	_ "github.com/lib/pq"
	"fmt"
	"os"
	"time"
)

// singleton Postgres instance to use pooling
// use this var by calling GetDB()
var db *sql.DB

func GetDB() *sql.DB {
	return db
}

func ConnectDB() error {
	url := os.Getenv("DB_URL")

	var err error
	db, err = sql.Open("postgres", url)
	if err != nil {
		return fmt.Errorf("sql.Open error: %w", err)
	}

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(5 * time.Minute)

	if err := db.Ping(); err != nil {
		return fmt.Errorf("db.Ping error: %w", err)
	}

	return nil
}
