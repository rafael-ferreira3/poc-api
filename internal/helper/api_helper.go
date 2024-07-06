package helper

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"
)

func ReadRequestBody(r *http.Request, result interface{}) error {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(result)
	return err
}

func WriteJson(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

func StringToInt64(s string) (int64, error) {
	return strconv.ParseInt(s, 0, 64)
}

func GetAddress() string {
	addr := os.Getenv("API_PORT")
	if addr == "" {
		addr = "8081"
	}

	return ":" + addr
}
