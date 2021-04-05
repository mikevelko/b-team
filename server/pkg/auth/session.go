package auth

import (
	"context"
	"fmt"
	"net/http"

	"github.com/pw-software-engineering/b-team/server/pkg/bookly"
)

type sessionHeaderName string

const (
	// UserIDHeader is a header name used to store user id in authenticated request
	UserIDHeader sessionHeaderName = "X-User-ID"
	// todo: add similar headers for the rest of data
)

// SetSessionHeaders sets headers with data from session
func SetSessionHeaders(header http.Header, session *bookly.Session) {
	panic("implement me")
}

// getSessionFromHeaders retrieves session from request headers
func getSessionFromHeaders(header http.Header) (*bookly.Session, error) {
	panic("implement me")
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
