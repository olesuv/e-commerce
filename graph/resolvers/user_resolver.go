package resolvers

import (
	"server.go/models"
	"server.go/services"

	"server.go/graph/model"
)

type UserResolver struct {
	userService *services.UserService
}

func NewUserResolver() *UserResolver {
	userService := services.NewUserService()
	return &UserResolver{userService}
}

func (r *UserResolver) CreateUser(input *model.CreateUserInput) (*models.User, error) {
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

func (r *UserResolver) GetUser(email string) (*models.User, error) {
	user, err := r.userService.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserResolver) DeleteUser(email string) (*models.User, error) {
	user, err := r.userService.DeleteUserByEmail(email)
	if err != nil {
		return nil, err
	}

	return user, nil
}
