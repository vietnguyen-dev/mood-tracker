package main

import (
	"log"
	"os"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/vietnguyen-dev/go-server/routes"
	"github.com/vietnguyen-dev/go-server/utils"
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
	api := r.PathPrefix("/api").Subrouter()
	api.HandleFunc("/moods/{user_id}", routes.GetMoods).Methods("GET")
	api.HandleFunc("/moods/{user_id}", routes.InsertMood).Methods("POST")
	api.HandleFunc("/moods", routes.UpdateMood).Methods("PUT")
	api.HandleFunc("/moods", routes.DeleteMood).Methods("DELETE")
	http.ListenAndServe(":8080", r)
}


