package routes

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/vietnguyen-dev/go-server/routes/models"
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

	rows, err := db.Query("SELECT * FROM vw_reports where user_id = ?;", user_id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	var reports []models.ReportRequest
	for rows.Next() {
		var report models.ReportRequest
		err := rows.Scan(&report.Id, &report.UserId, &report.MoodData, &report.StartDate, &report.EndDate, &report.CreatedAt, &report.UpdatedAt, &report.DeletedAt)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		reports = append(reports, report)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(reports)
}
