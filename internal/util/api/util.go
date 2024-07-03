package api

import (
	"encoding/json"
	"net/http"
)

type ApiError struct {
	Error string `json:"error"`
}

func WriteJson(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}
