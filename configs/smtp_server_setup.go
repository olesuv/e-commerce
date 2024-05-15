package configs

import (
	"log"
	"os"
	"strconv"
)

type SMTPServerConfig struct {
	SMTPLink     string
	SMTPPort     int
	SMTPUsername string
	SMTPPassword string
}

func GetSMTPServerConfig() SMTPServerConfig {
	smtpLink := os.Getenv("SMTP_LINK")
	smtpPortStr := os.Getenv("SMTP_PORT")
	smtpUsername := os.Getenv("SMTP_USERNAME")
	smtpPassword := os.Getenv("SMTP_PASSWORD")

	if smtpLink == "" {
		log.Fatal("server: smtp link is required")
	}

	smtpPort, err := strconv.Atoi(smtpPortStr)
	if err != nil {
		log.Fatal("server: smtp port is required and must be a valid integer")
	}

	if smtpUsername == "" {
		log.Fatal("server: smtp username is required")
	}

	if smtpPassword == "" {
		log.Fatal("server: smtp password is required")
	}

	return SMTPServerConfig{
		SMTPLink:     smtpLink,
		SMTPPort:     smtpPort,
		SMTPUsername: smtpUsername,
		SMTPPassword: smtpPassword,
	}
}
