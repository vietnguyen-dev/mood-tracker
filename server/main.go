package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"database/sql"
	"log"
	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

type Mood struct {
    ID     int    `json:"id"`
    Mood int    `json:"mood"`
    Note   string `json:"note"`
	UserId int `json:"user_id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt sql.NullString `json:"updated_at"`
	DeletedAt sql.NullString `json:"deleted_at"`
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", GetMoodsHandler)
	http.ListenAndServe(":8080", r)
}

func GetMoodsHandler(w http.ResponseWriter, r *http.Request) {
    db, err := sql.Open("sqlite3", "/Users/vietnguyen/test.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	
	fmt.Println("Successfully connected to SQLite!")
	
	stmt, err := db.Prepare("SELECT * FROM moods;")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer stmt.Close()
    
    rows, err := stmt.Query()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    moods := []Mood{}
    for rows.Next() {
        var m Mood
        err := rows.Scan(&m.ID, &m.Mood, &m.Note, &m.UserId, &m.CreatedAt, &m.UpdatedAt, &m.DeletedAt)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        moods = append(moods, m)
    }

    // Check for errors from iterating over rows
    if err = rows.Err(); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    if err := json.NewEncoder(w).Encode(moods); err != nil {
        http.Error(w, "Error encoding response", http.StatusInternalServerError)
        return
    }
}
