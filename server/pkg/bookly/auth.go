package bookly

import "context"

// Session represents the data attached to user session
type Session struct { // todo: fill in needed fields, e.g. userID, user role, hotelID etc
}

// Token represents the session token which is set in the user's browser.
//	Might change that to json marshallable struct in the future
type Token string

// SessionVerifier verifies if token is valid and not expired
type SessionVerifier interface {
	Verify(ctx context.Context, token Token) (*Session, error)
}

// SessionStorage persists sessions.
type SessionStorage interface { // todo: think what's needed here and implement it in postgres. Use it in SessionVerifier
	// todo: nice to have: in-memory cache so that request latency is lower
}

// UserVerifier verifies user credentials
type UserVerifier interface {
	Verify(email, password string) (Token, error)
}

// Hasher can hash any string and compare a hash with password's string hash
type Hasher interface {
	Hash(password string) string
	Compare(hash, password string) error
}
