package graph

//go:generate go run github.com/99designs/gqlgen generate

import (
	"server.go/services"
)

type Resolver struct {
	userService *services.UserService
}

func NewResolver() *Resolver {
	userService := services.NewUserService()
	return &Resolver{userService: userService}
}
