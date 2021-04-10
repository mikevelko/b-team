package main

import (
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
		r.Route("/session", func(r chi.Router) {
			r.Post("/", a.handleAuthorize)
		})
	})
}

func (a *api) handleAuthorize(w http.ResponseWriter, r *http.Request) {
	// token := r.Header.Get("x-session-token")
	// todo: check token structure, define it, json.Unmarshall it and pass to real verifier
	// todo: fill Token structure
	session, err := a.verifier.Verify(r.Context(), bookly.Token{})
	if err != nil {
		if errors.Is(err, bookly.ErrUserNotAuthenticated) {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	auth.SetSessionHeaders(w.Header(), session)
	return
}
