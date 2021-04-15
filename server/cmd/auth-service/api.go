package main

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/pw-software-engineering/b-team/server/pkg/auth"
	"github.com/pw-software-engineering/b-team/server/pkg/bookly"

	"github.com/go-chi/chi"
	"go.uber.org/zap"
)

type api struct {
	logger   *zap.Logger
	verifier bookly.SessionVerifier
}

func newAPI(logger *zap.Logger, verifier bookly.SessionVerifier) *api {
	return &api{
		logger:   logger,
		verifier: verifier,
	}
}

func (a *api) mount(router chi.Router) {
	router.Route("/api/v1/", func(r chi.Router) {
		r.Route("/session/", func(r chi.Router) {
			r.Post("/client/", a.handleAuthorizeClient)
			r.Post("/hotel/", a.handleAuthorizeHotel)
		})
	})
}

const (
	// HeaderHotelToken is the name for a token received from hotel UI
	HeaderHotelToken = "x-hotel-token"
	// HeaderClientToken is the name for a token received from client UI
	HeaderClientToken = "x-session-token"
)

func (a *api) handleAuthorizeClient(w http.ResponseWriter, r *http.Request) {
	a.logger.Info("Handling client auth")
	JSONToken := r.Header.Get(HeaderClientToken)

	a.verifyToken(w, r, JSONToken)
}

func (a *api) handleAuthorizeHotel(w http.ResponseWriter, r *http.Request) {
	a.logger.Info("Handling hotel auth")
	JSONToken := r.Header.Get(HeaderHotelToken)

	a.verifyToken(w, r, JSONToken)
}

func (a *api) verifyToken(w http.ResponseWriter, r *http.Request, JSONToken string) {
	var token bookly.Token
	if err := json.Unmarshal([]byte(JSONToken), &token); err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		a.logger.Info("user unauthorized, due to error when unmarshalling", zap.Error(err))
		return
	}

	session, err := a.verifier.Verify(r.Context(), token)
	if err != nil {
		a.logger.Info("user was not authorized:", zap.Error(err))
		if errors.Is(err, bookly.ErrUserNotAuthenticated) {
			w.WriteHeader(http.StatusUnauthorized)
			a.logger.Info("user unauthorized, because session is invalid", zap.Error(err))
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		a.logger.Info("user unauthorized, because of internal error in verifier", zap.Error(err))
		return
	}
	auth.SetSessionHeader(w.Header(), session)
	return
}
