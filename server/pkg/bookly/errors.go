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

// ErrOfferNotOwned indicates that hotel manager tried to manage offer that doesn't belong to managed hotel
var ErrOfferNotOwned = errors.New("offer does not belong to active hotel")

// ErrOfferStillActive indicates that hotel manager tried to delete offer that is still active
var ErrOfferStillActive = errors.New("offer is still active")

// ErrOfferDeleted indicates that offer is deleted and should not be modified
var ErrOfferDeleted = errors.New("offer is marked as deleted")
