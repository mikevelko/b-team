package main

import (
	"context"

	"github.com/pw-software-engineering/b-team/server/pkg/bookly"
)

type stubSessionVerifier struct{}

var _ bookly.SessionVerifier = &stubSessionVerifier{}

// Verify is stub implementation of bookly.SessionVerifier
// it checks if token exists - if it does, it confirms user
func (s *stubSessionVerifier) Verify(ctx context.Context, token bookly.Token) (*bookly.Session, error) {
	if token.ID == 0 {
		return nil, bookly.ErrUserNotAuthenticated
	}
	return &bookly.Session{}, nil
}

// todo: create verifier with session storage which would check session validity with database
