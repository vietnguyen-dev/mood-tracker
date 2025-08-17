package utils

import (
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
    "log"
	"os"
	"github.com/joho/godotenv"
)

var DB *sql.DB

// Connect opens the SQLite database and assigns it to the global variable
func Connect() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbPath := os.Getenv("DB_PATH")
    var dbErr error
    DB, dbErr = sql.Open("sqlite3", dbPath)
    if dbErr != nil {
        log.Fatalf("Failed to open database: %v", err)
    }

    // Optional: test the connection
    if testErr := DB.Ping(); testErr != nil {
        log.Fatalf("Failed to ping database: %v", testErr)
    }

    log.Println("Database connected!")
}

// Close safely closes the database connection
func Close() {
    if DB != nil {
        if err := DB.Close(); err != nil {
            log.Printf("Failed to close database: %v", err)
        } else {
            log.Println("Database closed!")
        }
    }
}
