package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
	"github.com/vietnguyen-dev/go-server/middleware"
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

	r := mux.NewRouter()
	r.Use(middleware.Logging)
	r.Use(middleware.ApiKeyAuth)

	api := r.PathPrefix("/api").Subrouter()

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:5173"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type", "x-api-key", "Access-Control-Allow-Origin", "Application/json"},
	})
	handler := c.Handler(r)

	// Moods
	api.HandleFunc("/moods/{user_id}", routes.GetMoods).Methods("GET")
	api.HandleFunc("/moods/{user_id}", routes.InsertMood).Methods("POST")
	api.HandleFunc("/moods", routes.UpdateMood).Methods("PUT")
	api.HandleFunc("/moods", routes.DeleteMood).Methods("DELETE")

	api.HandleFunc("/openai/generate-report", routes.GenerateReport).Methods("GET")


	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", handler)
}
