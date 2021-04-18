package bookly

import "errors"

// ErrUserNotAuthenticated indicates that user was not authenticated properly
var ErrUserNotAuthenticated = errors.New("user was not authenticated")

// ErrSessionExpired indicates that user session has expired
var ErrSessionExpired = errors.New("this session has expired")

// ErrSessionNotFound indicates that users session is not found on server (perhaps of incorrect session data)
var ErrSessionNotFound = errors.New("session not found")

// ErrEmptyHotelName indicates that hotel manager tried to change hotel name to empty string
var ErrEmptyHotelName = errors.New("hotel name is empty")
