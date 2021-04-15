package main

import (
	"context"
	"fmt"
	"regexp"

	"github.com/pw-software-engineering/b-team/server/pkg/bookly"
)

type userService struct {
	userStorage    bookly.UserStorage
	sessionStorage bookly.SessionStorage
}

func newUserService(userStorage bookly.UserStorage, sessionStorage bookly.SessionStorage) *userService {
	return &userService{userStorage: userStorage, sessionStorage: sessionStorage}
}

// UserVerify handles verify user
func (u *userService) UserVerify(ctx context.Context, userName string, password string) (bool, bookly.Token, error) {
	err := IsUserNameValid(userName)
	if err != nil {
		return false, bookly.Token{}, err
	}

	check, user, errU := u.userStorage.UserVerify(ctx, userName, password)
	if errU != nil {
		return false, bookly.Token{}, errU
	}
	if check == false {
		return false, bookly.Token{}, nil
	}

	token, errS := u.sessionStorage.CreateNew(ctx, user.ID)
	if errS != nil {
		return false, bookly.Token{}, errS
	}

	return true, token, nil
}

// GetUser handles getting user info
func (u *userService) GetUser(ctx context.Context, userID int64) (bookly.User, error) {
	// todo: Get userID from SessionMiddleware header
	user, err := u.userStorage.GetUser(ctx, userID)
	if err != nil {
		return bookly.User{}, err
	}
	return user, nil
}

// UpdateUserInformation handles update user info
func (u *userService) UpdateUserInformation(ctx context.Context, userName string, email string, userID int64) error {
	err := IsUserNameValid(userName)
	if err != nil {
		return err
	}

	err = IsEmailValid(email)
	if err != nil {
		return err
	}

	err = u.userStorage.UpdateUserInformation(ctx, userID, userName, email)
	if err != nil {
		return err
	}
	return nil
}

var _ bookly.UserService = &userService{}

// IsUserNameValid check if UserName is valid
func IsUserNameValid(userName string) error {
	if len(userName) > bookly.UserNameMaxLen || len(userName) < bookly.UserNameMinLen {
		return fmt.Errorf("userName: wrong length")
	}
	return nil
}

// IsEmailValid check if Email is valid
func IsEmailValid(email string) error {
	if len(email) > bookly.EmailMaxLen || len(email) < bookly.EmailMinLen {
		return fmt.Errorf("email: wrong length")
	}

	re := regexp.MustCompile(bookly.EmailValidationRegexp)

	if false == re.MatchString(email) {
		return fmt.Errorf("email: wrong formatt")
	}
	return nil
}
