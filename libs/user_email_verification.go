package libs

import (
	"context"
	"fmt"

	"github.com/go-gomail/gomail"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

func GenerateVerificationToken(ctx context.Context, email string, rdb *redis.Client) error {
	token := uuid.New().String()

	err := rdb.Set(ctx, token, email, 0).Err()
	if err != nil {
		return fmt.Errorf("server: redis set, details: %w", err)
	}

	return nil
}

func SendVerificationEmail(email string, token string) {
	m := gomail.NewMessage()
	m.SetHeader("From", "your@email.com")
	m.SetHeader("To", email)
	m.SetHeader("Subject", "Email Verification")
	m.SetBody("text/plain", "Please click on this link to verify your email: http://yourwebsite.com/verifyuser="+token)

	d := gomail.NewDialer("smtp.yourserver.com", 587, "your_username", "your_password")

	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}
