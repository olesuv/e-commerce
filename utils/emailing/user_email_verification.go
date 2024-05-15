package utils

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/go-gomail/gomail"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"server.go/configs"
	"server.go/constants"
)

func SendVerificationEmail(ctx context.Context, emailUser string, rdb *redis.Client) error {
	emailSender := os.Getenv("EMAIL_SENDER")
	if emailSender == "" {
		panic("server: email sender is required")
	}

	smtpConfig := configs.GetSMTPServerConfig()

	d := gomail.NewDialer(
		smtpConfig.SMTPLink,
		smtpConfig.SMTPPort,
		smtpConfig.SMTPUsername,
		smtpConfig.SMTPPassword,
	)

	token, err := insertVerificationToken(context.Background(), emailUser, rdb)
	if err != nil {
		return err
	}

	m := createVerificationEmailMessage(emailSender, emailUser, token)

	if err := d.DialAndSend(m); err != nil {
		return err
	}

	log.Println("server: email sent to ", emailUser)
	return nil
}

func insertVerificationToken(ctx context.Context, email string, rdb *redis.Client) (string, error) {
	token := uuid.New().String()

	err := rdb.Set(ctx, token, email, 0).Err()
	if err != nil {
		return "", fmt.Errorf("server: redis set, details: %w", err)
	}

	return token, nil
}

func createVerificationEmailMessage(emailSender string, emailUser string, token string) *gomail.Message {
	m := gomail.NewMessage()
	m.SetHeader("From", emailSender)
	m.SetHeader("To", emailUser)
	m.SetHeader("Subject", constants.EMAIL_HEADER)
	m.SetBody("text/plain", constants.EMAIL_BODY+constants.VERIFICATION_LINK+token)
	return m
}
