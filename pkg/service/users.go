package service

import (
	"fmt"

	"github.com/DimKush/guestbook/tree/main/internal/entities/User"
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

func (data *UsersServiceWorker) GetUserByUsername(username string) (User.User, error) {
	user, err := data.db_users.GetUserByUsername(username)

	if err != nil {
		return User.User{}, err
	}
	if (user == User.User{}) {
		return User.User{}, fmt.Errorf("User with username = %s doesn't exists.", username)
	}

	return user, nil
}

func (data *UsersServiceWorker) GetUsersByParams(filter *User.User) ([]User.User, error) {
	// users, err := data.db_users.GetUsersByParams(filter)

	// if err != nil {
	// 	log.Error().Msg(err.Error())
	// }
	return nil, nil
}

func InitUsersServiceWorker(repos repository.UsersService) *UsersServiceWorker {
	return &UsersServiceWorker{
		db_users: repos,
	}
}
