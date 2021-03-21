package httputils

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	error string
}

func RespondWithError(w http.ResponseWriter, msg string) {
	errResponse := ErrorResponse{msg}
	js, err := json.Marshal(errResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(js)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
