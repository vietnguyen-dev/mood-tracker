//go:build ignore
// +build ignore

package utils

import (
	"crypto/rand"
	"encoding/base64"
)

func CreateAPIKey() string {
	bytes := make([]byte, 32)
	_, err := rand.Read(bytes)
	if err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(bytes)
}