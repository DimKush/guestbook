package service

import (
	"net/smtp"
	"os"

	"github.com/DimKush/guestbook/tree/main/internal/entities/EmailEventDb"
	"github.com/DimKush/guestbook/tree/main/pkg/repository"
)

type EmailServiceAuth struct {
	event_db repository.EmailService
}

func (data *EmailServiceAuth) InitEmailEvent(email_event EmailEventDb.EmailEventDb) error {
	from := os.Getenv("")
	password := os.Getenv("")

	to := []string{
		"demonmadman228@gmail.com",
		"sillyuglyseal@gmail.com",
	}
	// smtp server config
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Message
	message := []byte("Hello there")

	auth := smtp.PlainAuth("", from, password, smtpHost)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	return err
}

func InitEmailService(repos repository.EmailService) *EmailServiceAuth {
	return &EmailServiceAuth{
		event_db: repos,
	}
}
