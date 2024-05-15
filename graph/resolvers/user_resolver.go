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

	var userName string
	if input.Name == nil || *input.Name == "" {
		userName = "New User"
	} else {
		userName = *input.Name
	}

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
	// example of how to get cookie from context
	// if userEmail := middleware.CtxValue(ctx); userEmail == "" {
	// 	return nil, fmt.Errorf("login first")
	// }

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
	if *input.Email == "" || *input.Password == "" {
		return "", fmt.Errorf("email and password are required")
	}

	user, err := r.userService.GetUserByEmail(*input.Email)
	if err != nil && err.Error() == "mongo: no documents in result" {
		return "", fmt.Errorf("user not found")
	}
	if err != nil {
		return "", fmt.Errorf("server: get user by email, details: %w", err)
	}

	userHash := user.Password
	comapre := authHelpers.VerifyPassword(*input.Password, userHash)

	if !comapre {
		return "", fmt.Errorf("invalid password")
	}

	token, err := authHelpers.GenearteJwtToken(ctx, *input.Email)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (r *UserResolver) Orders(ctx context.Context, obj *models.User) ([]*models.Order, error) {
	panic(fmt.Errorf("not implemented: Orders - orders"))
}
