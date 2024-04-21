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
