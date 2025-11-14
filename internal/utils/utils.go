package utils

import (
	"crypto/sha256"
	"encoding/json"
	"net/http"
)

func HashPassword(password string) string {
	hash := sha256.New()
	hash.Write([]byte(password))
	return string(hash.Sum(nil))
}

func WriteJSON(w http.ResponseWriter, data interface{}, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}
