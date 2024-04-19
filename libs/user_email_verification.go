package libs

import (
	"github.com/go-gomail/gomail"
	"github.com/google/uuid"
)

func GenerateVerificationToken(email string) (string, error) {
	token := uuid.New().String()
	return token, nil
}

func SendVerificationEmail(email string) {
	m := gomail.NewMessage()
	m.SetHeader("From", "your@email.com")
	m.SetHeader("To", email)
	m.SetHeader("Subject", "Email Verification")
	m.SetBody("text/plain", "Please click on this link to verify your email: http://yourwebsite.com/verify?email="+email)

	d := gomail.NewDialer("smtp.yourserver.com", 587, "your_username", "your_password")

	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}
