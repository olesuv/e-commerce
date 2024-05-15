package utils

import (
	"fmt"
	"net/mail"

	"server.go/graph/model"
	"server.go/models"
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

func (ue *UserErrors) CheckLoginUserInput(input model.LoginUserInput) (models.User, error) {
	if *input.Email == "" || *input.Password == "" {
		return models.User{}, fmt.Errorf("email and password are required")
	}

	user, err := ue.userService.GetUserByEmail(*input.Email)
	if err != nil && err.Error() == "mongo: no documents in result" {
		return models.User{}, fmt.Errorf("user not found")
	}
	if err != nil {
		return models.User{}, fmt.Errorf("server: get user by email, details: %w", err)
	}

	return *user, nil
}

func (ue *UserErrors) SetName(input *model.CreateUserInput) string {
	var userName string

	if input.Name == nil || *input.Name == "" {
		userName = "New User"
	} else {
		userName = *input.Name
	}

	return userName
}
