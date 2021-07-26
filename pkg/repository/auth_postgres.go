package repository

import (
	"fmt"

	"github.com/DimKush/guestbook/tree/main/internal/entities/User"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

type AuthPostgres struct {
	db *gorm.DB
}

func InitAuthPostgres(database *gorm.DB) *AuthPostgres {
	return &AuthPostgres{db: database}
}

func (data *AuthPostgres) CreateUser(user User.User) (int, error) {
	var count int64
	data.db.Table(users).Where("username=?", user.Username).Count(&count)
	if count != 0 {
		return 0, fmt.Errorf("Username with username : %s already exists.", user.Username)
	}

	err := data.db.Table(users).Create(&user).Error

	if err != nil {
		log.Error().Msgf("Error during execute the query : \n%s.", err.Error())
		return 0, err
	}

	return user.Id, nil
}
