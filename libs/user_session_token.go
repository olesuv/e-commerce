package libs

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"server.go/models"
	"server.go/services"
)

func GenearteJwtToken(email string, uas *services.UserAuthService) (string, error) {
	claims := jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(time.Hour * 24 * 40).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		return "", fmt.Errorf("libs: jwt secret is required")
	}

	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", fmt.Errorf("libs: signed string, details: %w", err)
	}

	_, err = uas.CreateSession(&models.Session{
		UserEmail:   email,
		DateExpired: time.Now().Add(time.Hour * 24 * 40).String(),
		Token:       tokenString,
	})
	if err != nil {
		return "", fmt.Errorf("server: create session, details: %w", err)
	}

	return tokenString, nil
}

func ValidateJwtToken(tokenString string, uas *services.UserAuthService) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		jwtSecret := os.Getenv("JWT_SECRET")
		if jwtSecret == "" {
			return nil, fmt.Errorf("libs: jwt secret is required")
		}
		return []byte(jwtSecret), nil
	})
	if err != nil {
		return "", fmt.Errorf("libs: parse token, details: %w", err)
	}
	if !token.Valid {
		return "", fmt.Errorf("libs: invalid token")
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", fmt.Errorf("libs: invalid claims")
	}
	email, ok := claims["email"].(string)
	if !ok {
		return "", fmt.Errorf("libs: invalid email claim")
	}
	session, err := uas.GetSessionByToken(tokenString)
	if err != nil {
		return "", fmt.Errorf("server: get session by token, details: %w", err)
	}
	if session == nil || session.UserEmail != email {
		return "", fmt.Errorf("server: invalid session")
	}
	return email, nil
}
