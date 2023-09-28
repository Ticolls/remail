package email

import (
	"crypto/tls"
	"os"

	"github.com/joho/godotenv"
	"gopkg.in/gomail.v2"
)

var (
	mail   *gomail.Message
	dialer *gomail.Dialer
)

func Init() error {
	err := godotenv.Load()

	if err != nil {
		return err
	}

	mail = gomail.NewMessage()

	senderEmail := os.Getenv("SENDER_EMAIL")
	senderPassword := os.Getenv("SENDER_PASSWORD")
	receiverEmail := os.Getenv("RECEIVER_EMAIL")

	// Set E-Mail sender
	mail.SetHeader("From", senderEmail)

	// Set E-Mail receivers
	mail.SetHeader("To", receiverEmail)

	// Settings for SMTP server
	dialer = gomail.NewDialer("smtp-mail.outlook.com", 587, senderEmail, senderPassword)

	// This is only needed when SSL/TLS certificate is not valid on server.
	// In production this should be set to false.
	dialer.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	return nil
}
