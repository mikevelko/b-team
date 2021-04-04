package httputils

import "net/http"

//WriteJSONResponse writes json object to http reply
func WriteJSONResponse(w http.ResponseWriter, js []byte) {
	w.Header().Set("Content-Type", "application/json")
	_, err := w.Write(js)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
