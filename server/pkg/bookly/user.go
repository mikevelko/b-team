package bookly

import (
	"context"

	"gopkg.in/guregu/null.v3"
)

// predefined string for user roles
const (
	UserRoleHotelManager      = "HOTEL_MANAGER"
	UserRoleHotelStaff        = "HOTEL_STAFF"
	UserRoleHotelReceptionist = "HOTEL_RECEPTIONIST"
	UserRoleClientCustomer    = "CLIENT_CUSTOMER"
)

// predefined const for user verification
const (
	EmailValidationRegexp = "^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$"
	EmailMinLen           = 3
	EmailMaxLen           = 254
	UserNameMinLen        = 3
	UserNameMaxLen        = 254
)

// User is a domain-level model
type User struct {
	ID        int64
	FirstName string
	Surname   string
	Email     string
	UserName  string
	UserRole  string
	// Password                    string
	HotelID null.Int
}

// UserStorage persists user related data
type UserStorage interface {
	AddUserForce(ctx context.Context, user User, password string) (int64, error)
	GetUser(ctx context.Context, userID int64) (User, error)
	UpdateUserInformation(ctx context.Context, id int64, userName string, email string) error
	UserVerify(ctx context.Context, userName string, password string) (bool, User, error)
}

// UserService is a service which is responsible for actions related to user
type UserService interface {
	UserVerify(ctx context.Context, userName string, password string) (bool, Token, error)
	GetUser(ctx context.Context, userID int64) (User, error)
	UpdateUserInformation(ctx context.Context, userName string, email string, userID int64) error
}
