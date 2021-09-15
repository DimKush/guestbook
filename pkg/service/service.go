package service

import (
	"github.com/DimKush/guestbook/tree/main/internal/entities/EmailEventDb"
	"github.com/DimKush/guestbook/tree/main/internal/entities/Item"
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

type EmailService interface {
	InitEmailEvent(email_event EmailEventDb.EmailEventDb) error
}

type ListService interface {
	GetAllLists() ([]List.List, error)
	GetListsByParams(List.List) ([]List.List, error)
	GetListById(list_id int) (List.List, error)
	GetAutoListId() (int, error)
	CreateList(List.List) error
	DeleteListById(list_id int) error
	UpdateListById(list *List.List) error
}

type ItemsService interface {
	GetItemsByParams(Item.Item) ([]Item.Item, error)
	CreateNewItem(Item.Item) error
	GetItemTypesByParams(Item.ItemType) ([]Item.ItemType, error)
	GetItemsAvailability(int, UserIn.UserIn) (int, error)
}

type UsersSevice interface {
	GetAllUsernames() ([]string, error)
	GetUserByUsername(username string) (User.User, error)
}

type Service struct {
	Authorization
	Event
	EmailService
	ListService
	UsersSevice
	ItemsService
}

func ServiceInit(repos *repository.Repository) *Service {
	return &Service{
		Authorization: InitAuthService(repos.Authorization, repos.UsersService, repos.EmailService),
		ListService:   InitListsServiceWorker(repos.ListService, repos.UsersService),
		UsersSevice:   InitUsersServiceWorker(repos.UsersService),
		ItemsService:  InitItemsServiceWorker(repos.ItemsService),
	}
}
