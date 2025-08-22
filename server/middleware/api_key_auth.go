package middleware

import (
	"net/http"
	"github.com/vietnguyen-dev/go-server/utils"
	"crypto/sha256"
	"encoding/base64"
	"database/sql"
	"log"
)

type ApiKeyModel struct {
	ID        int64     `json:"id"`
	APIKey    string    `json:"hashed_key"`
}

func ApiKeyAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apiKey := r.Header.Get("x-api-key")
		if apiKey == "" {
			http.Error(w, "API key is required", http.StatusUnauthorized)
			return
		}
		db := utils.GetDB()
		
		// hash the key
        hashed := sha256.Sum256(apiKey)
        hashedBytes := hashed[:]  // convert array -> slice
		byteToString := base64.StdEncoding.EncodeToString(hashedBytes)

		var id int64
		var hashedKey string
		err := db.QueryRow("SELECT * FROM vw_api_keys WHERE hashed_key = ?;", hashedBytes).Scan(&id, &hashedKey)
		if sql.ErrNoRows == err {
			http.Error(w, "API key is not registered", http.StatusUnauthorized)
			log.Println("Error querying API key: ", err, "hashedKey: ", hashedKey, "byteToString: ", byteToString)
			return
		}
		if err != nil {
			http.Error(w, "Error querying API key", http.StatusUnauthorized)
			log.Println("Error querying API key: ", err, "hashedKey: ", hashedKey, "byteToString: ", byteToString)
			return
		}
		if hashedKey == "" {
			http.Error(w, "Invalid API key", http.StatusUnauthorized)
			return
		}

		// if the key is valid, continue to the next handler
		if hashedKey == byteToString {
			next.ServeHTTP(w, r)
		}
	})
}