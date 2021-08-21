package service

import (
	"github.com/DimKush/guestbook/tree/main/internal/entities/EmailEventDb"
	"github.com/DimKush/guestbook/tree/main/internal/entities/List"
	"github.com/DimKush/guestbook/tree/main/internal/entities/User"
	"github.com/DimKush/guestbook/tree/main/internal/entities/UserIn"
	"github.com/DimKush/guestbook/tree/main/pkg/repository"
)

type Authorization interface {
	CreateUser(user User.User) (int, error)
	CheckUserExitstsWithPass(user UserIn.UserIn) error
	CheckUserExitsts(user UserIn.UserIn) error
	GenerateToken(username, password string) (string, error)
	ParseToken(accessToken string) (int, string, error)
	GetUser(userIn UserIn.UserIn) (User.User, error)
}

type Event interface {
}

type EventList interface {
}

type EmailService interface {
	InitEmailEvent(email_event EmailEventDb.EmailEventDb) error
}

type ListService interface {
	GetAllLists() ([]List.List, error)
	GetListsByParams(List.List) ([]List.List, error)
}

type UsersSevice interface {
	GetAllUsernames() ([]string, error)
}

type Service struct {
	Authorization
	Event
	EventList
	EmailService
	ListService
	UsersSevice
}

func ServiceInit(repos *repository.Repository) *Service {
	return &Service{
		Authorization: InitAuthService(repos.Authorization, repos.EmailService),
		ListService:   InitListsServiceWorker(repos.ListService),
		UsersSevice:   InitUsersServiceWorker(repos.UsersService),
	}
}
