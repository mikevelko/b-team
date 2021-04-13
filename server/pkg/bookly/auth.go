package bookly

import "context"

// Session represents the data attached to user session
type Session struct {
	UserID  int64 `json:"x-user-id"`
	HotelID int64 `json:"x-hotel-id"`
}

// Token represents the session token which is set in the user's browser.
type Token struct {
	ID        int64  `json:"id"`
	CreatedAt string `json:"createdAt"`
}

// SessionVerifier verifies if token is valid and not expired
type SessionVerifier interface {
	Verify(ctx context.Context, token Token) (*Session, error)
}

// SessionStorage persists sessions.
type SessionStorage interface {
	CreateNew(ctx context.Context, userID int64) (Token, error)
	Update(ctx context.Context, token Token) error
	GetSession(ctx context.Context, token Token) (*Session, error)
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
