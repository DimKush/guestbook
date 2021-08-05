package repository

import (
	"github.com/DimKush/guestbook/tree/main/internal/entities/AuditEvent"
	"github.com/DimKush/guestbook/tree/main/internal/entities/EmailEventDb"
	"github.com/DimKush/guestbook/tree/main/internal/entities/User"
	"gorm.io/gorm"
)

type Authorization interface {
	CreateUser(user User.User) (int, error)
	GetUser(username, password string) (User.User, error)
}

type Event interface {
}

type EventList interface {
}

type EmailEvent interface {
	CreateEmailEvent(email_event EmailEventDb.EmailEventDb) error
}

type AuditInt interface {
	WriteEvent(AuditEvent.AuditEvent) error
}

type Repository struct {
	Authorization
	Event
	EventList
	AuditInt
	EmailEvent
}

func RepositoryInit(db *gorm.DB) *Repository {
	return &Repository{
		Authorization: InitAuthPostgres(db),
		AuditInt:      InitAuditRep(db),
		EmailEvent:    InitEmailEvent(db),
	}
}
