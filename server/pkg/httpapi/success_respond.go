package httpapi

import "net/http"

// RespondWithCode respond with code and set content type application/json
func RespondWithCode(w http.ResponseWriter, code int) {
	// Specification require response content type set as application/json
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)
}
