package httpapi

import (
	"encoding/json"
	"net/http"
)

// ErrorResponse is json-serializable response body from http api
type ErrorResponse struct {
	Error string `json:"error"`
}

// RespondWithError is used to handle internal errors in httpapi
func RespondWithError(w http.ResponseWriter, msg string) {
	errResponse := ErrorResponse{msg}
	js, err := json.Marshal(errResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusBadRequest)
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(js)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
