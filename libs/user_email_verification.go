package libs

import (
	"github.com/go-gomail/gomail"
	"github.com/google/uuid"
)

func GenerateVerificationToken(email string) (string, error) {
	token := uuid.New().String()
	// implement redis logic
	return token, nil
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
