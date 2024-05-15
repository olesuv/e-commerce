package resolvers

import (
	"context"
	"fmt"
	"net/mail"

	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"server.go/graph/model"
	"server.go/models"
	"server.go/services"
	"server.go/utils"
	emailing "server.go/utils/emailing"
)

type UserResolver struct {
	userService *services.UserService
	rdb         *redis.Client
}

func NewUserResolver(rdb *redis.Client) *UserResolver {
	return &UserResolver{
		userService: services.NewUserService(),
		rdb:         rdb,
	}
}

func (r *UserResolver) CreateUser(ctx context.Context, input model.CreateUserInput) (*models.User, error) {
	if input.Email == nil || *input.Email == "" {
		return nil, fmt.Errorf("email is required")
	}

	if input.Password == nil || *input.Password == "" {
		return nil, fmt.Errorf("password is required")
	}

	email := *input.Email

	_, err := mail.ParseAddress(email)
	if err != nil {
		return nil, fmt.Errorf("invalid email address")
	}

	user, err := r.userService.GetUserByEmail(*input.Email)

	if err == nil && user != nil {
		return nil, fmt.Errorf("user with this email already exists")
	}
	if err != nil && err.Error() != "mongo: no documents in result" {
		return nil, fmt.Errorf("server: get user by email, details: %w", err)
	}

	if err != nil && err.Error() == "mongo: no documents in result" {
		defaultName := "New User"
		if input.Name == nil || *input.Name == "" {
			input.Name = &defaultName
		}

		hashedPassword := utils.HashPassword(*input.Password)

		user, err = r.userService.CreateUser(&models.User{
			Id:       primitive.NewObjectID(),
			Name:     *input.Name,
			Email:    *input.Email,
			Password: hashedPassword,
		})
		if err != nil {
			return nil, fmt.Errorf("server: create user, details: %w", err)
		}

		verificationToken, err := emailing.GenerateVerificationToken(ctx, *input.Email, r.rdb)
		if err != nil {
			_, err = r.userService.DeleteUserByEmail(*input.Email)
			if err != nil {
				return nil, fmt.Errorf("server: delete user by email, details: %w", err)
			}

			return nil, fmt.Errorf("server: generate verification token, details: %w", err)
		}

		err = emailing.SendVerificationEmail(*input.Email, verificationToken)
		if err != nil {
			_, err := r.userService.DeleteUserByEmail(*input.Email)
			if err != nil {
				return nil, fmt.Errorf("server: delete user by email, details: %w", err)
			}

			// FIX: Google email sending error
			return nil, fmt.Errorf("server: send verification email, details: %w", err)
		}
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
	comapre := utils.VerifyPassword(*input.Password, userHash)

	if !comapre {
		return "", fmt.Errorf("invalid password")
	}

	token, err := utils.GenearteJwtToken(ctx, *input.Email)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (r *UserResolver) Orders(ctx context.Context, obj *models.User) ([]*models.Order, error) {
	panic(fmt.Errorf("not implemented: Orders - orders"))
}
