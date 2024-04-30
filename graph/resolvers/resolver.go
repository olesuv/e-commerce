package resolvers

import "github.com/redis/go-redis/v9"

//go:generate go run github.com/99designs/gqlgen generate

type Resolver struct {
	*UserResolver
}

func NewResolver() *Resolver {
	return &Resolver{
		UserResolver: NewUserResolver(&redis.Client{}),
	}
}
