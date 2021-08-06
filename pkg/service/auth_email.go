package service

import (
	"fmt"
	"net/smtp"
	"os"

	"github.com/DimKush/guestbook/tree/main/internal/entities/EmailEventDb"
	"github.com/DimKush/guestbook/tree/main/pkg/repository"
)

type EmailServiceAuth struct {
	event_db repository.EmailService
}

func (data *EmailServiceAuth) InitEmailEvent(email_event *EmailEventDb.EmailEventDb) error {
	fmt.Println("HERE!")
	fmt.Println(email_event)

	email_event.Sender = os.Getenv("EMAIL_SENDER")
	email_event.SenderPass = os.Getenv("EMAIL_PASSWORD")

	// smtp server config
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Message
	message := []byte("Hello there")

	auth := smtp.PlainAuth("", email_event.Sender, email_event.SenderPass, smtpHost)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, email_event.Sender, []string{email_event.Receiver}, message)
	return err
}

func InitEmailService(repos repository.EmailService) *EmailServiceAuth {
	return &EmailServiceAuth{
		event_db: repos,
	}
}
