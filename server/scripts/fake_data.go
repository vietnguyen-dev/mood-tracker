//go:build ignore
// +build ignore

package main

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

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
	user_id := os.Getenv("TEST_USER_ID")
	fmt.Println("dbPath: ", dbPath)
	fmt.Println("user_id: ", user_id)
    var dbErr error
    DB, dbErr = sql.Open("sqlite3", dbPath)
    if dbErr != nil {
        log.Fatal("Error opening database")
    }
	defer DB.Close()
	// Array of 3 realistic mood notes
	moodNotes := []string{
		"Feeling productive and energetic today!",
		"Had a challenging day but staying positive.",
		"Great mood, everything is going well.",
	}
	
	i := 366
	
	for i >= 0 {
		random_mood := rand.Intn(10) + 1
		random_note := moodNotes[rand.Intn(len(moodNotes))]
		date_subtract := time.Now().AddDate(0, 0, -i)
		insert_date := date_subtract.Format("2006-01-02 15:04:05")
		_, err := DB.Exec("INSERT INTO moods (mood, notes, user_id, created_at) VALUES (?, ?, ?, ?);", random_mood, random_note, user_id, insert_date)
		if err != nil {
			log.Fatal("Error inserting data: ", err)
		}
		i--

		fmt.Println("mood: ", random_mood, ":notes: ", random_note, "user_id: ", user_id, "insert_date: ", insert_date)
	}
}