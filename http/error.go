package http

import (
	"encoding/json"
	"net/http"
)

type errorResponse struct {
	Err string `json:"error,omitempty"`
}

func invalidRequest(w http.ResponseWriter, r *http.Request, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(&errorResponse{"invalid request"})
}
