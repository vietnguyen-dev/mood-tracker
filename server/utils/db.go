package utils

import (
	"database/sql"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

// InitDB initializes the database connection pool once
func InitDB() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}

	dbPath := os.Getenv("DB_PATH")
	var dbErr error
	DB, dbErr = sql.Open("sqlite3", dbPath)
	if dbErr != nil {
		return dbErr
	}

	// Configure connection pool settings
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)
	DB.SetConnMaxLifetime(time.Hour)

	// Test the connection
	if testErr := DB.Ping(); testErr != nil {
		return testErr
	}

	log.Println("Database Pool Initialized!")
	return nil
}

// GetDB returns the database connection pool
func GetDB() *sql.DB {
	return DB
}

// CloseDB closes the entire connection pool (only call this when shutting down)
func CloseDB() error {
	if DB != nil {
		log.Println("Closing database pool...")
		return DB.Close()
	}
	return nil
}
