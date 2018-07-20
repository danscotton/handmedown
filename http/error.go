package http

import (
	"encoding/json"
	"net/http"
)

type errorResponse struct {
	Err string `json:"error,omitempty"`
}

func Error(w http.ResponseWriter, err string, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(&errorResponse{err})
}

func InvalidRequestError(w http.ResponseWriter) {
	Error(w, "invalid request", http.StatusBadRequest)
}
