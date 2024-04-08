package graph

//go:generate go run github.com/99designs/gqlgen generate

import (
	"context"

	"server.go/graph/model"
	"server.go/models"
	"server.go/services"
)

type Resolver struct {
	userService *services.UserService
}

func ServicesResolver() *Resolver {
	userService := services.NewUserService()
	return &Resolver{userService}
}

func (r *Resolver) CreateUser(ctx context.Context, input model.CreateUserInput) (*models.User, error) {
	user, err := r.userService.CreateUser(&models.User{
		Name:  *input.Name,
		Email: input.Email,
		Phone: *input.Phone,
	})

	if err != nil {
		return nil, err
	}

	return user, nil
}
