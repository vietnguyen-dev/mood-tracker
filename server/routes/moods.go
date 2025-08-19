package routes

import (
	"net/http"
	"github.com/vietnguyen-dev/go-server/utils"
	"github.com/gorilla/mux"
	"database/sql"
	"encoding/json"
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

func MoodsHandler(w http.ResponseWriter, r *http.Request) {
	utils.Connect()
	defer utils.Close()

	vars := mux.Vars(r)	
	user_id := vars["user_id"]

	rows, err := utils.DB.Query("SELECT * FROM vw_moods where user_id = ?;", user_id)
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