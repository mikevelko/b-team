package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pw-software-engineering/b-team/server/pkg/bookly"
)

const (
	// HeaderXSession is a header name. The header is used to store session data
	HeaderXSession = "X-Session"
)

// SetSessionHeader sets headers with data from session
func SetSessionHeader(header http.Header, session *bookly.Session) {
	encoded, err := json.Marshal(session)
	if err != nil {
		panic(fmt.Errorf("error marshaling session: %w", err))
	}
	header.Set(HeaderXSession, string(encoded))
}

// getSessionFromHeaders retrieves session from request headers
func getSessionFromHeaders(header http.Header) (*bookly.Session, error) {
	encoded := header.Get(HeaderXSession)
	var session bookly.Session
	if err := json.Unmarshal([]byte(encoded), &session); err != nil {
		return nil, fmt.Errorf("auth: could not unmarshall header")
	}
	return &session, nil
}

// SessionMiddleware creates middleware which adds session data to context. Use SessionFromContext to retrieve it back.
func SessionMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			session, err := getSessionFromHeaders(r.Header)
			if err != nil {
				http.Error(w, "internal auth error", http.StatusInternalServerError)
				return
			}
			r = r.WithContext(sessionToContext(r.Context(), session))
			next.ServeHTTP(w, r)
		})
	}
}

type ctxUserKey struct{}

// sessionToContext adds user to given context. User SessionFromContext to retrieve it back.
func sessionToContext(ctx context.Context, user *bookly.Session) context.Context {
	return context.WithValue(ctx, ctxUserKey{}, user)
}

// SessionFromContext finds the user from the context. REQUIRES SessionMiddleware to have run.
func SessionFromContext(ctx context.Context) *bookly.Session {
	user, ok := ctx.Value(ctxUserKey{}).(*bookly.Session)
	if !ok {
		// this is allowable, because it would never happen under normal condition.
		// in worst case scenario recoverer middleware would catch that panic and 500 status would be returned
		panic(fmt.Sprintf("could not retrieve session from context, got %T. Is SessionMiddleware added?", ctx.Value(ctxUserKey{})))
	}
	return user
}
