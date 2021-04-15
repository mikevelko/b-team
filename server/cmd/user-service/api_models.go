package main

import "github.com/pw-software-engineering/b-team/server/pkg/bookly"

// GetUserResponse respond struct
type GetUserResponse struct {
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

// GetUserResponseFromUser func
func GetUserResponseFromUser(user bookly.User) GetUserResponse {
	respond := GetUserResponse{
		Name:     user.FirstName,
		Surname:  user.Surname,
		Username: user.UserName,
		Email:    user.Email,
	}
	return respond
}

// PatchUserRequest respond struct
type PatchUserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

// PathUserErrorResponse error struct
type PathUserErrorResponse struct {
	EmailError    ErrorResponse `json:"Invalid-email-format"`
	UserNameError ErrorResponse `json:"Invalid-username"`
}

// ErrorResponse error struct
type ErrorResponse struct {
	Message string `json:"errorDescription"`
}

// PostUserRequest request struct
type PostUserRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

// PostUserErrorResponse error struct
type PostUserErrorResponse struct {
	Desc string `json:"desc"`
}
