package routes

import (
	"net/http"
	"github.com/vietnguyen-dev/go-server/utils"
	"github.com/gorilla/mux"
	"database/sql"
	"encoding/json"
	"io"
	"fmt"
	"time"
)

type Mood struct {
	ID int `json:"id"`
	Mood int `json:"mood"`
	Note string `json:"note"`
	UserId int `json:"user_id"`
	CreatedAt sql.NullString `json:"created_at"`
	UpdatedAt sql.NullString `json:"updated_at"`
	DeletedAt sql.NullString `json:"deleted_at"`
}

func GetMoods(w http.ResponseWriter, r *http.Request) {
	// Get database connection from pool
	db := utils.GetDB()
	
	vars := mux.Vars(r)	
	user_id := vars["user_id"]
	start_date := r.URL.Query().Get("start_date")
	end_date := r.URL.Query().Get("end_date")

	rows, err := db.Query("SELECT * FROM vw_moods where user_id = ? AND created_at >= ? AND created_at <= ?;", user_id, start_date, end_date)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var moods []Mood
	for rows.Next() {
		var mood Mood
		err := rows.Scan(&mood.ID, &mood.Mood, &mood.Note, &mood.UserId, &mood.CreatedAt, &mood.UpdatedAt, &mood.DeletedAt)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		moods = append(moods, mood)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(moods)
}

type MoodRequest struct {
	Mood int `json:"mood"`
	Note string `json:"note"`
}

func InsertMood(w http.ResponseWriter, r *http.Request) {
	db := utils.GetDB()
	vars := mux.Vars(r)
	user_id := vars["user_id"]
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var moodRequest MoodRequest
	err = json.Unmarshal(body, &moodRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	insert, err := db.Exec("INSERT INTO moods (mood, notes, user_id, created_at) VALUES (?, ?, ?, ?);", moodRequest.Mood, moodRequest.Note, user_id, time.Now().Format("2006-01-02 15:04:05"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	insert_id, err := insert.LastInsertId()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("Mood inserted successfully with id: %d", insert_id)))
}

type EditMoodRequest struct {
	ID int `json:"id"`
	Mood int `json:"mood"`
	Notes string `json:"notes"`
}

func UpdateMood(w http.ResponseWriter, r *http.Request) {
	db := utils.GetDB()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var editMoodRequest EditMoodRequest
	err = json.Unmarshal(body, &editMoodRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	edit, err := db.Exec("UPDATE moods SET mood = ?, notes = ?, updated_at = ? WHERE id = ?;", editMoodRequest.Mood, editMoodRequest.Notes, time.Now().Format("2006-01-02 15:04:05"), editMoodRequest.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	rows_affected, err := edit.RowsAffected()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Mood updated successfully with rows affected: %d", rows_affected)))
}

func DeleteMood(w http.ResponseWriter, r *http.Request) {
	db := utils.GetDB()
	delete_id:= r.URL.Query().Get("delete_id")
	if delete_id == "" {
		http.Error(w, "Missing delete_id", http.StatusBadRequest)
		return
	}

	deleted, err := db.Exec("UPDATE moods SET deleted_at = ? WHERE id = ?;", time.Now().Format("2006-01-02 15:04:05"), delete_id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	rows_affected, err := deleted.RowsAffected()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Mood deleted successfully with rows affected: %d", rows_affected)))
}