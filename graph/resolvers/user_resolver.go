package resolvers

import (
	"context"

	"server.go/graph/model"
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

func (r *UserResolver) DeleteUser(ctx context.Context, email string) (*models.User, error) {
	user, err := r.userService.DeleteUserByEmail(email)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserResolver) Users(ctx context.Context) ([]*models.User, error) {
	users, err := r.userService.GetUsers()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (r *UserResolver) ID(ctx context.Context, obj *models.User) (string, error) {
	return string(obj.Id.Hex()), nil
}

func (r *UserResolver) User(ctx context.Context, email string) (*models.User, error) {
	user, err := r.userService.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}

	return user, nil
}
