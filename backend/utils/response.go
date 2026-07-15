package utils

import (
	"encoding/json"
	"net/http"
)

func JSONSuccess(w http.ResponseWriter, status int, message string, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	json.NewEncoder(w).Encode(map[string]any{
		"success": true,
		"message": message,
		"data":    data,
	})
}

func JSONError(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	json.NewEncoder(w).Encode(map[string]any{
		"success": false,
		"message": message,
	})
}
