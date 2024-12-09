package utils

import (
	"encoding/json"
	"net/http"
)

func DecodeBody(r *http.Request, v any) error {
	return json.NewDecoder(r.Body).Decode(v)
}

func WriteJSON(w http.ResponseWriter, payload any, code int) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	return json.NewEncoder(w).Encode(payload)
}

func WriteError(w http.ResponseWriter, msg string, code int) error {
	return WriteJSON(w, map[string]string{"error": msg}, code)
}

func WriteMessage(w http.ResponseWriter, msg string, code int) error {
	return WriteJSON(w, map[string]string{"message": msg}, code)
}
