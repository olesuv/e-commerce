package utils

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/go-gomail/gomail"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"server.go/constants"
)

func GenerateVerificationToken(ctx context.Context, email string, rdb *redis.Client) (string, error) {
	token := uuid.New().String()

	err := rdb.Set(ctx, token, email, 0).Err()
	if err != nil {
		return "", fmt.Errorf("server: redis set, details: %w", err)
	}

	return token, nil
}

func SendVerificationEmail(emailUser string, token string) error {
	emailSender := os.Getenv("EMAIL_SENDER")
	if emailSender == "" {
		panic("server: email sender is required")
	}

	m := gomail.NewMessage()
	m.SetHeader("From", emailSender)
	m.SetHeader("To", emailUser)
	m.SetHeader("Subject", constants.EMAIL_HEADER)
	m.SetBody("text/plain", constants.EMAIL_BODY+constants.VERIFICATION_LINK+token)

	smtpLink := os.Getenv("SMTP_LINK")
	smtpPortStr := os.Getenv("SMTP_PORT")
	smtpUsername := os.Getenv("SMTP_USERNAME")
	smtpPassword := os.Getenv("SMTP_PASSWORD")

	if smtpLink == "" {
		log.Fatal("server: smtp link is required")
	}
	smtpPort, err := strconv.Atoi(smtpPortStr)
	if err != nil {
		log.Fatal("server: smtp port is required")
	}
	if smtpUsername == "" {
		log.Fatal("server: smtp username is required")
	}
	if smtpPassword == "" {
		log.Fatal("server: smtp password is required")
	}

	d := gomail.NewDialer(smtpLink, smtpPort, smtpUsername, smtpPassword)

	if err := d.DialAndSend(m); err != nil {
		return err
	}

	log.Println("server: email sent to ", emailUser)
	return nil
}
