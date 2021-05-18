package httpapi

import (
	"encoding/json"
	"net/http"

	"go.uber.org/zap"
)

// WriteJSONResponse writes json object to http reply
func WriteJSONResponse(logger *zap.Logger, w http.ResponseWriter, marshallable interface{}) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(marshallable); err != nil {
		logger.Error("handlePostOffer: could not marshal response", zap.Error(err))
		RespondWithError(w, "could not encode response JSON")
	}
}

// JSONRespondError respond with error code and json struct
func JSONRespondError(w http.ResponseWriter, err interface{}, code int) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	errJ := json.NewEncoder(w).Encode(err)
	if errJ != nil {
		return errJ
	}

	return nil
}
