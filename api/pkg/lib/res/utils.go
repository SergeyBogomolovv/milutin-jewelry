package res

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func WriteJSON(w http.ResponseWriter, payload any, code int) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	return json.NewEncoder(w).Encode(payload)
}

func WriteError(w http.ResponseWriter, msg string, status int) {
	WriteJSON(w, ErrorResponse{Error: msg}, status)
}

func WriteMessage(w http.ResponseWriter, msg string, status int) {
	WriteJSON(w, MessageResponse{Message: msg}, status)
}

func DecodeBody(r *http.Request, v any) error {
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(v); err != nil {
		return err
	}
	if err := decoder.Decode(&struct{}{}); !errors.Is(err, io.EOF) {
		return errors.New("body must contain a single JSON value")
	}
	return nil
}
