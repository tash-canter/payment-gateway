package services

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func setupDatabase() error {
	var err error
	db, err = sql.Open("sqlite3", "payments.db")
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	query := `
	CREATE TABLE IF NOT EXISTS payments (
		id TEXT PRIMARY KEY,
		card_number TEXT,
		expiry_month INTEGER,
		expiry_year INTEGER,
		amount INTEGER,
		currency TEXT,
		status TEXT,
		processed_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	)`
	_, err = db.Exec(query)
	if err != nil {
		return fmt.Errorf("failed to create table: %w", err)
	}

	return nil
}

func init() {
	err := setupDatabase()
	if err != nil {
		log.Fatalf("Database initialization failed: %v", err)
	}
}
