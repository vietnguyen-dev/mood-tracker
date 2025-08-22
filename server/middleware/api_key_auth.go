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
		decodedKey, err := base64.StdEncoding.DecodeString(apiKey)
        if err != nil {
            http.Error(w, "Invalid API key", http.StatusUnauthorized)
            return
        }

        hashed := sha256.Sum256(decodedKey)
        hashedBase64 := base64.StdEncoding.EncodeToString(hashed[:])

        var id int64
        var storedHash string
        err = db.QueryRow("SELECT id, hashed_key FROM api_keys WHERE hashed_key = ?", hashedBase64).Scan(&id, &storedHash)
        if err == sql.ErrNoRows {
            http.Error(w, "API key is not registered", http.StatusUnauthorized)
            return
        }
        if err != nil {
            http.Error(w, "Error querying API key", http.StatusInternalServerError)
            log.Println("DB error:", err)
            return
        }

        next.ServeHTTP(w, r)
	})
}