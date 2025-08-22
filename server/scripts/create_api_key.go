//go:build ignore
// +build ignore

package main

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"log"
	"os"
	"fmt"
	"database/sql"
	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	dbPath := os.Getenv("DB_PATH")
	fmt.Println("dbPath: ", dbPath)
    var dbErr error
    DB, dbErr = sql.Open("sqlite3", dbPath)
	defer DB.Close()
    if dbErr != nil {
        log.Fatal("Error opening database")
    }

	generatedKey := make([]byte, 32) // Generate 32 random bytes

	// Read random bytes into the slice
	_, ReadErr := rand.Read(generatedKey)
	if ReadErr != nil {
		log.Fatalf("Error generating random bytes: %v", ReadErr)
	}

	fmt.Printf("Generated random bytes: %x\n", generatedKey)
	clientKey := base64.StdEncoding.EncodeToString(generatedKey)
	hashedKey := sha256.Sum256(generatedKey)
	hashedKeyToStore := base64.StdEncoding.EncodeToString(hashedKey[:])
	fmt.Printf("Hashed key: %x\n", hashedKeyToStore)

	insert, err := DB.Exec("INSERT INTO api_keys (hashed_key) VALUES (?);", hashedKeyToStore)
	if err != nil {
		log.Fatalf("Error inserting API key: %v", err)
	}
	fmt.Printf("For Client: %v\n", clientKey)
	fmt.Printf("Inserted: %v\n", insert)
}