package http

import (
	"encoding/json"
	"net/http"
)

type response struct {
	w http.ResponseWriter
}

func (r *response) With(code int, body interface{}) {
	r.w.Header().Set("Content-Type", "application/json")
	r.w.WriteHeader(code)
	json.NewEncoder(r.w).Encode(body)
}

func Respond(w http.ResponseWriter) *response {
	return &response{w: w}
}
