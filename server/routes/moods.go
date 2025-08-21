package routes

import (
	"net/http"
	"github.com/vietnguyen-dev/go-server/utils"
	"github.com/gorilla/mux"
	"encoding/json"
	"fmt"
	"time"
	"github.com/vietnguyen-dev/go-server/routes/models"
)

func GetMoods(w http.ResponseWriter, r *http.Request) {
	// Get database connection from pool
	db := utils.GetDB()
	user_id := mux.Vars(r)["user_id"]
	if user_id == "" {
		http.Error(w, "no user id", http.StatusBadRequest)
		return
	}
	start_date := r.URL.Query().Get("start_date")
	end_date := r.URL.Query().Get("end_date")
	if start_date == "" || end_date == "" {
		http.Error(w, "start_date and end_date are required", http.StatusBadRequest)
		return
	}

	rows, err := db.Query("SELECT * FROM vw_moods where user_id = ? AND created_at >= ? AND created_at <= ?;", user_id, start_date, end_date)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var moods []models.Mood
	for rows.Next() {
		var mood models.Mood
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

func InsertMood(w http.ResponseWriter, r *http.Request) {
	var moodRequest models.MoodRequest
	err := json.NewDecoder(r.Body).Decode(&moodRequest)
	if err := moodRequest.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db := utils.GetDB()
	vars := mux.Vars(r)
	user_id := vars["user_id"]
	
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

func UpdateMood(w http.ResponseWriter, r *http.Request) {
	var editMoodRequest models.EditMoodRequest
	err := json.NewDecoder(r.Body).Decode(&editMoodRequest)
	if err := editMoodRequest.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db := utils.GetDB()
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