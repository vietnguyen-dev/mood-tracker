package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/vietnguyen-dev/go-server/utils"
)

func GetReports(w http.ResponseWriter, r *http.Request) {
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
}

