package routes

import (
	"encoding/json"
	"fmt"
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

	page := r.URL.Query().Get("page")
	amount := r.URL.Query().Get("amount")
	if page == "" {
		page = "1"
	}
	if amount == "" {
		amount = "10"
	}

	query := fmt.Sprintf("SELECT * FROM vw_reports where user_id = %s LIMIT %s OFFSET (%s * %s);", user_id, amount, page, amount)

	rows, err := db.Query(query)
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
