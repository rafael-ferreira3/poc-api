package util

import (
	"encoding/json"
	"net/http"
	"os"
)

type ApiError struct {
	Error string `json:"error"`
}

func WriteJson(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

func GetAddress() string {
	addr := os.Getenv("API_PORT")
	if addr == "" {
		addr = "8081"
	}

	return ":" + addr
}
