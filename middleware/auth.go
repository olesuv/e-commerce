package middleware

import (
	"context"
	"fmt"
	"net/http"

	"server.go/services"
)

type AuthTokenKey string

type AuthMiddleware struct {
	UserAuthService *services.UserAuthService
}

func NewAuthMiddleware(uvs *services.UserAuthService) *AuthMiddleware {
	return &AuthMiddleware{UserAuthService: uvs}
}

func (am *AuthMiddleware) Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		isValid, err := am.isValidToken(token)
		if err != nil || !isValid {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), AuthTokenKey("token"), token)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}

func (am *AuthMiddleware) isValidToken(token string) (bool, error) {
	session, err := am.UserAuthService.GetSessionByToken(token)
	if err != nil {
		return false, fmt.Errorf("middleware: get session by token, details: %w", err)
	}

	if session.Token != token {
		return false, fmt.Errorf("middleware: invalid token")
	}

	return true, nil
}
