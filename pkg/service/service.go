package service

import (
	"github.com/DimKush/guestbook/tree/main/internal/entities/User"
	"github.com/DimKush/guestbook/tree/main/pkg/repository"
)

type Authorization interface {
	CreateUser(user User.User) (int, error)
}

type Event interface {
}

type EventList interface {
}

type Service struct {
	Authorization
	Event
	EventList
}

func ServiceInit(repos *repository.Repository) *Service {
	return &Service{
		Authorization: InitAuthService(repos.Authorization),
	}
}
