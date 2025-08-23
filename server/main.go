package main

import (
	"log"
	"os"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/vietnguyen-dev/go-server/routes"
	"github.com/vietnguyen-dev/go-server/utils"
	"github.com/vietnguyen-dev/go-server/middleware"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	
	// Initialize database connection pool
	if err := utils.InitDB(); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer utils.CloseDB()
	
	host := os.Getenv("HOST")
	r := mux.NewRouter()
	r.Host(host)
	r.Use(middleware.Logging)
	r.Use(middleware.ApiKeyAuth)
	api := r.PathPrefix("/api").Subrouter()

	// Moods
	api.HandleFunc("/moods/{user_id}", routes.GetMoods).Methods("GET")
	api.HandleFunc("/moods/{user_id}", routes.InsertMood).Methods("POST")
	api.HandleFunc("/moods", routes.UpdateMood).Methods("PUT")
	api.HandleFunc("/moods", routes.DeleteMood).Methods("DELETE")

	// Reports
	api.HandleFunc("/reports", routes.GetReports).Methods("GET")
	api.HandleFunc("/reports", routes.InsertReport).Methods("POST")
	api.HandleFunc("/reports", routes.UpdateReport).Methods("PUT")
	api.HandleFunc("/reports", routes.DeleteReport).Methods("DELETE")

	http.ListenAndServe(":8080", r)
}


