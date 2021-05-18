package httpapi

import "net/http"

// IsHeaderTypeValid checks if http request header has proper header and send error response if not
func IsHeaderTypeValid(w http.ResponseWriter, r *http.Request, expectedType string, errMsg string) bool {
	contentType := r.Header.Get("Content-Type")
	if contentType != expectedType {
		RespondWithError(w, errMsg)
		return false
	}
	return true
}
