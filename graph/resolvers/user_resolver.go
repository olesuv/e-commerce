package resolvers

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"server.go/graph/model"
	"server.go/libs"
	"server.go/models"
	"server.go/services"
)

type UserResolver struct {
	userService *services.UserService
}

func NewUserResolver() *UserResolver {
	return &UserResolver{services.NewUserService()}
}

func (r *UserResolver) CreateUser(ctx context.Context, input model.CreateUserInput) (*models.User, error) {
	if input.Name == nil || *input.Email == "" {
		return nil, fmt.Errorf("email is required")
	}

	user, err := r.userService.GetUserByEmail(*input.Email)
	if err != nil && err.Error() != "mongo: no documents in result" {
		return nil, fmt.Errorf("server: get user by email, details: %w", err)
	}
	if err != nil && err.Error() == "mongo: no documents in result" {
		defaultName := "New User"
		if input.Name == nil || *input.Name == "" {
			input.Name = &defaultName
		}

		hashedPassword := libs.HashPassword(*input.Password)

		user, err = r.userService.CreateUser(&models.User{
			Id:       primitive.NewObjectID(),
			Name:     *input.Name,
			Email:    *input.Email,
			Phone:    *input.Phone,
			Password: hashedPassword,
			Image:    *input.Image,
		})
		if err != nil {
			return nil, fmt.Errorf("server: create user details: %w", err)
		}
	}

	return user, nil
}

func (r *UserResolver) DeleteUser(ctx context.Context, email string) (*models.User, error) {
	user, err := r.userService.DeleteUserByEmail(email)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserResolver) ID(ctx context.Context, obj *models.User) (string, error) {
	return obj.Id.Hex(), nil
}

func (r *UserResolver) Users(ctx context.Context) ([]*models.User, error) {
	users, err := r.userService.GetUsers()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (r *UserResolver) User(ctx context.Context, email string) (*models.User, error) {
	if email == "" {
		return nil, fmt.Errorf("server: email is required")
	}

	user, err := r.userService.GetUserByEmail(email)
	if user == nil {
		return nil, fmt.Errorf("server: email already exists")
	}
	if err != nil {
		return nil, err
	}

	return user, nil
}

// Orders is the resolver for the orders field.
func (r *UserResolver) Orders(ctx context.Context, obj *models.User) ([]*models.Order, error) {
	panic(fmt.Errorf("not implemented: Orders - orders"))
}
