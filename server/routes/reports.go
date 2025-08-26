package routes

import (
	"fmt"
	"net/http"
)

func GetReports(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetReports")
}