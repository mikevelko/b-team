package bookly

import "context"

// Session represents the data attached to user session
type Session struct {
	UserID     int
	HotelToken string
	Token      Token
}

// Token represents the session token which is set in the user's browser.
//	Might change that to json marshallable struct in the future
type Token struct {
	ID        int    `json:"id"`
	CreatedAt string `json:"createdAt"`
}

// SessionVerifier verifies if token is valid and not expired
type SessionVerifier interface {
	Verify(ctx context.Context, token Token) (*Session, error)
}

// SessionStorage persists sessions.
type SessionStorage interface {
	CreateNew(username string, password string) (Token, error)
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
