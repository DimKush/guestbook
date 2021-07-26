package repository

import (
	"github.com/DimKush/guestbook/tree/main/internal/entities/User"
	"gorm.io/gorm"
)

type Authorization interface {
	CreateUser(user User.User) (int, error)
}

type Event interface {
}

type EventList interface {
}

type Repository struct {
	Authorization
	Event
	EventList
}

func RepositoryInit(db *gorm.DB) *Repository {
	return &Repository{
		Authorization: InitAuthPostgres(db),
	}
}
