package service

import (
	"github.com/DimKush/guestbook/tree/main/internal/entities/EmailEventDb"
	"github.com/DimKush/guestbook/tree/main/internal/entities/User"
	"github.com/DimKush/guestbook/tree/main/internal/entities/UserIn"
	"github.com/DimKush/guestbook/tree/main/pkg/repository"
)

type Authorization interface {
	CreateUser(user User.User) (int, error)
	CheckUserExitsts(user UserIn.UserIn) error
	GenerateToken(username, password string) (string, error)
	ParseToken(accessToken string) (int, error)
}

type Event interface {
}

type EventList interface {
}

type EmailService interface {
	InitEmailEvent(email_event EmailEventDb.EmailEventDb) error
}

type Service struct {
	Authorization
	Event
	EventList
	EmailService
}

func ServiceInit(repos *repository.Repository) *Service {
	return &Service{
		Authorization: InitAuthService(repos.Authorization, repos.EmailService),
	}
}
