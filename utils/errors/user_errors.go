package utils

import (
	"fmt"
	"net/mail"

	"server.go/graph/model"
	"server.go/services"
)

type UserErrors struct {
	userService *services.UserService
}

func NewUserErrors() *UserErrors {
	return &UserErrors{
		userService: services.NewUserService(),
	}
}

func (ue *UserErrors) CheckCreateUserInput(input model.CreateUserInput) error {
	if input.Email == nil || *input.Email == "" {
		return fmt.Errorf("email is required")
	}

	if input.Password == nil || *input.Password == "" {
		return fmt.Errorf("password is required")
	}

	email := *input.Email

	_, err := mail.ParseAddress(email)
	if err != nil {
		return fmt.Errorf("invalid email address")
	}

	user, err := ue.userService.GetUserByEmail(*input.Email)
	if err != nil {
		if err.Error() != "mongo: no documents in result" {
			return fmt.Errorf("server: get user by email, details: %w", err)
		}
		return nil
	}
	if user != nil {
		return fmt.Errorf("user with this email already exists")
	}

	return nil
}
