package handlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/redis/go-redis/v9"
	"server.go/services"
)

type UserEmailVerificationHandler struct {
	userService *services.UserService
	rdb         *redis.Client
}

func NewUserEmailVerificationHandler(rdb *redis.Client) *UserEmailVerificationHandler {
	return &UserEmailVerificationHandler{
		userService: services.NewUserService(),
		rdb:         rdb,
	}
}

func (evh *UserEmailVerificationHandler) VerifyUserEmailHandler(rw http.ResponseWriter, req *http.Request) {
	token := req.URL.Query().Get("token")
	if token == "" {
		http.Error(rw, "redis: token is required", http.StatusBadRequest)
		return
	}

	email, err := evh.rdb.Get(context.Background(), token).Result()
	// TODO: make normal error check if user already activated email
	if err != nil {
		http.Error(rw, fmt.Sprintf("Link already used.\nDetails: %v", err), http.StatusInternalServerError)
		return
	}
	if email == "" {
		http.Error(rw, "redis: invalid token", http.StatusBadRequest)
		return
	}

	_, err = evh.userService.VerifyUserEmail(email)
	if err != nil {
		http.Error(rw, fmt.Sprintf("redis: verifying user by email: %v", err), http.StatusInternalServerError)
		return
	}

	err = evh.rdb.Del(context.Background(), token).Err()
	if err != nil {
		http.Error(rw, fmt.Sprintf("redis: error deleting token from Redis: %v", err), http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte("User verified successfully"))
}
