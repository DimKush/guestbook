package service

import (
	"github.com/DimKush/guestbook/tree/main/pkg/repository"
)

type UsersServiceWorker struct {
	db_users repository.UsersService
}

func (data *UsersServiceWorker) GetAllUsernames() ([]string, error) {
	users, err := data.db_users.GetAllUsernames()
	var usernames []string
	if err != nil {
		return nil, err
	}

	for _, val := range users {
		usernames = append(usernames, val.Username)
	}

	return usernames, nil
}

func InitUsersServiceWorker(repos repository.UsersService) *UsersServiceWorker {
	return &UsersServiceWorker{
		db_users: repos,
	}
}
