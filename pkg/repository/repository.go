package repository

import (
	"github.com/DimKush/guestbook/tree/main/internal/entities/AuditEvent"
	"github.com/DimKush/guestbook/tree/main/internal/entities/EmailEventDb"
	"github.com/DimKush/guestbook/tree/main/internal/entities/EventItem"
	"github.com/DimKush/guestbook/tree/main/internal/entities/List"
	"github.com/DimKush/guestbook/tree/main/internal/entities/User"
	"github.com/DimKush/guestbook/tree/main/internal/entities/UserIn"
	"gorm.io/gorm"
)

type Authorization interface {
	CreateUser(user User.User) (int, error)
	GetUser(username, password string) (User.User, error)
	GetUserByUserIn(userIn UserIn.UserIn) (User.User, error)
}

type Event interface {
}

type EventList interface {
}

type EmailService interface {
	InitEmailEvent(email_event EmailEventDb.EmailEventDb) error
}

type AuditInt interface {
	WriteEvent(AuditEvent.AuditEvent) error
}

type ListService interface {
	GetAllLists() ([]List.List, error)
	GetListsByParams(List.List) ([]List.List, error)
	GetListById(list_id int) (List.List, error)
	GetAutoListId() (int, error)
	CreateList(List.List) error
	DeleteListById(list_id int) error
}

type EventsService interface {
	GetEventsByParams(EventItem.EventItem) ([]EventItem.EventItem, error)
}

type UsersService interface {
	GetAllUsernames() ([]UserIn.UserIn, error)
	GetUserByUsername(username string) (User.User, error)
}

type Repository struct {
	Authorization
	Event
	EventList
	AuditInt
	EmailService
	ListService
	EventsService
	UsersService
}

func RepositoryInit(db *gorm.DB) *Repository {
	return &Repository{
		Authorization: InitAuthPostgres(db),
		AuditInt:      InitAuditRep(db),
		EmailService:  InitEmailEventRep(db),
		ListService:   InitListsRep(db),
		EventsService: InitEventsRep(db),
		UsersService:  InitUsersRepos(db),
	}
}
