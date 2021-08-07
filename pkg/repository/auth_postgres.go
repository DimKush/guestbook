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

	// check if user with username exists
	if err := data.db.Table(users).Where("username=?", user.Username).Count(&count).Error; err != nil {
		return 0, fmt.Errorf("SQL : Cannot select from table %s: Reason : %s", users, err.Error())
	}
	if count != 0 {
		return 0, fmt.Errorf("SQL : Username with username : %s already exists.", user.Username)
	}

	// check if user with email exists
	if err := data.db.Table(users).Where("email=?", user.Email).Count(&count).Error; err != nil {
		return 0, fmt.Errorf("SQL : Cannot select from table %s: Reason : %s", users, err.Error())
	}
	if count != 0 {
		return 0, fmt.Errorf("SQL : Email with email : %s already exists.", user.Email)
	}

	err := data.db.Table(users).Create(&user).Error

	if err != nil {
		log.Error().Msgf("SQL : Error during execute the query : \n%s.", err.Error())
		return 0, err
	}

	return user.Id, nil
}

func (data *AuthPostgres) GetUser(username, password string) (User.User, error) {
	var user User.User
	if err := data.db.Table(users).Where("username=?", username).Scan(&user).Error; err != nil {
		return User.User{}, fmt.Errorf("SQL : Cannot select from table %s: Reason : %s", users, err.Error())
	}
	if user.Username == "" {
		return User.User{}, fmt.Errorf("SQL : No rows in result set with username : %s", username)
	}

	return user, nil
}

func (data *AuthPostgres) GetUserByUsername(username string) (User.User, error) {
	var user User.User
	if err := data.db.Table(users).Where("username=?", username).Scan(&user).Error; err != nil {
		return User.User{}, fmt.Errorf("SQL : Cannot select from table %s: Reason : %s", users, err.Error())
	}
	if user.Username == "" {
		return User.User{}, fmt.Errorf("SQL : No rows in result set with username : %s", username)
	}

	return user, nil
}
