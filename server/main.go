package main

import (
	"log"
	"os"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/vietnguyen-dev/go-server/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	host := os.Getenv("HOST")
	r := mux.NewRouter()
	r.Host(host)
	r.HandleFunc("/api/moods/{user_id}", routes.MoodsHandler)
	http.ListenAndServe(":8080", r)
}


