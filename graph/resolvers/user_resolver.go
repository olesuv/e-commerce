package resolvers

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"server.go/graph/model"
	"server.go/models"
	"server.go/services"
	authHelpers "server.go/utils/auth"
	emailing "server.go/utils/emailing"
	errors "server.go/utils/errors"
)

type UserResolver struct {
	userService *services.UserService
	rdb         *redis.Client
	userErrors  *errors.UserErrors
}

func NewUserResolver(rdb *redis.Client) *UserResolver {
	return &UserResolver{
		userService: services.NewUserService(),
		userErrors:  errors.NewUserErrors(),
		rdb:         rdb,
	}
}

func (r *UserResolver) CreateUser(ctx context.Context, input model.CreateUserInput) (*models.User, error) {
	if err := r.userErrors.CheckCreateUserInput(input); err != nil {
		return nil, err
	}

	userName := r.userErrors.SetName(&input)
	hashedPassword := authHelpers.HashPassword(*input.Password)

	newUser := &models.User{
		Id:       primitive.NewObjectID(),
		Name:     userName,
		Email:    *input.Email,
		Password: hashedPassword,
	}

	user, err := r.userService.CreateUser(newUser)
	if err != nil {
		return nil, fmt.Errorf("server: create user, details: %w", err)
	}

	err = emailing.SendVerificationEmail(ctx, *input.Email, r.rdb)
	if err != nil {
		_, err := r.userService.DeleteUserByEmail(*input.Email)
		if err != nil {
			return nil, fmt.Errorf("server: delete user by email, details: %w", err)
		}

		// FIX: Google email sending error
		return nil, fmt.Errorf("server: send verification email, details: %w", err)
	}

	return user, nil
}

func (r *UserResolver) DeleteUser(ctx context.Context, email string) (*models.User, error) {
	if email == "" {
		return nil, fmt.Errorf("server: email is required")
	}

	user, err := r.userService.DeleteUserByEmail(email)

	if err.Error() == "mongo: no documents in result" {
		return nil, fmt.Errorf("server: no user with specified email")
	}

	if err != nil {
		return nil, fmt.Errorf("server: delete user by email, details: %w", err)
	}

	return user, nil
}

func (r *UserResolver) Users(ctx context.Context) ([]*models.User, error) {
	users, err := r.userService.GetUsers()
	if err != nil {
		return nil, fmt.Errorf("server: get users, details: %w", err)
	}

	return users, nil
}

func (r *UserResolver) ID(ctx context.Context, obj *models.User) (string, error) {
	return obj.Id.Hex(), nil
}

func (r *UserResolver) User(ctx context.Context, email string) (*models.User, error) {
	if email == "" {
		return nil, fmt.Errorf("server: email is required")
	}

	user, err := r.userService.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserResolver) LoginUser(ctx context.Context, input model.LoginUserInput) (string, error) {
	user, err := r.userErrors.CheckLoginUserInput(input)
	if err != nil {
		return "", err
	}

	comparePasswords, err := authHelpers.VerifyPassword(*input.Password, user.Password)
	if err != nil && !comparePasswords {
		return "", err
	}

	token, err := authHelpers.GenearteJwtToken(ctx, *input.Email)
	if err != nil {
		return "", err
	}

	return token, nil
}
