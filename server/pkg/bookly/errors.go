package bookly

import "errors"

// ErrUserNotAuthenticated indicates that user was not authenticated properly
var ErrUserNotAuthenticated = errors.New("user was not authenticated")
