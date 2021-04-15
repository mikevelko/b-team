package testutils

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

// DumpMiddleware func
func DumpMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := httputil.DumpRequest(r, true)
		fmt.Print(string(body))
		next.ServeHTTP(w, r)
	})
}
