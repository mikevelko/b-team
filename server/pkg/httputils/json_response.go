package httputils

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
