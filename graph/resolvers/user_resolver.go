package resolvers

import (
	"context"
	"fmt"
	"log"
	"net/mail"
	"os"

	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"server.go/configs"
	"server.go/graph/model"
	"server.go/libs"
	"server.go/middleware"
	"server.go/models"
	"server.go/services"
)

type UserResolver struct {
	userService *services.UserService
	rdb         *redis.Client
}

func NewUserResolver() *UserResolver {
	builder := configs.NewRedisClientBuilder()

	builder.WithAddr("redis:" + os.Getenv("REDIS_PORT")).WithPassword(os.Getenv("REDIS_PASSWORD"))

	rdb, err := builder.Build()
	if err != nil {
		log.Fatal("server: redis connection failed, details: ", err)
	}
	// if rdb.Ping(context.Background()).Err() != nil {
	// 	log.Fatal("server: redis ping failed")
	// }

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

	_, err := mail.ParseAddress(*input.Email)
	if err != nil {
		return nil, fmt.Errorf("invalid email address")
	}

	user, err := r.userService.GetUserByEmail(*input.Email)

	if err == nil && user != nil {
		return nil, fmt.Errorf("email already exists")
	}
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
			Password: hashedPassword,
		})
		if err != nil {
			return nil, fmt.Errorf("server: create user, details: %w", err)
		}

		verificationToken, err := libs.GenerateVerificationToken(ctx, *input.Email, r.rdb)
		if err != nil {
			// TODO: handle delete user possible error
			r.userService.DeleteUserByEmail(*input.Email)

			return nil, fmt.Errorf(err.Error())
		}

		libs.SendVerificationEmail(*input.Email, verificationToken)
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
	if userEmail := middleware.CtxValue(ctx); userEmail == "" {
		return nil, fmt.Errorf("login first")
	}

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

func (r *UserResolver) VerifyUser(ctx context.Context, token string) (*models.User, error) {
	if token == "" {
		return nil, fmt.Errorf("server: token is required")
	}

	email, err := r.rdb.Get(ctx, token).Result()
	if err != nil {
		return nil, fmt.Errorf("server: verify user by email, details: %w", err)
	}
	if email == "" {
		return nil, fmt.Errorf("server: invalid token")
	}

	user, err := r.userService.VerifyUser(email)
	if err != nil {
		return nil, fmt.Errorf("server: verify user by email, details: %w", err)
	}

	r.rdb.Del(ctx, token)

	_, err = libs.GenearteJwtToken(ctx, email)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	return user, nil
}

func (r *UserResolver) LoginUser(ctx context.Context, input model.LoginUserInput) (string, error) {
	if *input.Email == "" || *input.Password == "" {
		return "", fmt.Errorf("email and password are required")
	}

	user, err := r.userService.GetUserByEmail(*input.Email)
	if err != nil {
		return "", fmt.Errorf("server: get user by email, details: %w", err)
	}

	if user == nil {
		return "", fmt.Errorf("user not found")
	}

	userHash := user.Password
	comapre := libs.VerifyPassword(*input.Password, userHash)

	if !comapre {
		return "", fmt.Errorf("invalid password")
	}

	token, err := libs.GenearteJwtToken(ctx, *input.Email)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (r *UserResolver) Orders(ctx context.Context, obj *models.User) ([]*models.Order, error) {
	panic(fmt.Errorf("not implemented: Orders - orders"))
}
