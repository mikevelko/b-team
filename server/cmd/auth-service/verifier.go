package main

import (
	"context"

	"github.com/pw-software-engineering/b-team/server/pkg/bookly"
)

// SessionVerifier implements logic for verifying user sessions
type SessionVerifier struct {
	storage bookly.SessionStorage
}

var _ bookly.SessionVerifier = &SessionVerifier{}

// NewSessionVerifier is a constructor for SessionVerifier
func NewSessionVerifier(sessionStorage bookly.SessionStorage) *SessionVerifier {
	verifier := &SessionVerifier{
		storage: sessionStorage,
	}
	return verifier
}

// Verify checks if token exists - if it does, it confirms user
func (s *SessionVerifier) Verify(ctx context.Context, token bookly.Token) (*bookly.Session, error) {
	err := s.storage.Update(ctx, token)
	if err != nil {
		if err == bookly.ErrSessionNotFound {
			return nil, bookly.ErrUserNotAuthenticated
		}
		return nil, err
	}
	session, errSession := s.storage.GetSession(ctx, token)
	if errSession != nil {
		return nil, err
	}
	return session, nil
}
