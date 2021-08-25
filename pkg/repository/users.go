package repository

import (
	"fmt"

	"github.com/DimKush/guestbook/tree/main/internal/entities/User"
	"github.com/DimKush/guestbook/tree/main/internal/entities/UserIn"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

type UsersReposWorker struct {
	db gorm.DB
}

func (data *UsersReposWorker) GetAllUsernames() ([]UserIn.UserIn, error) {
	rows, err := data.db.Table(users).Select("username").Rows()

	var dbUsers []UserIn.UserIn
	if err != nil {
		log.Info().Msg(err.Error())
	}

	for rows.Next() {
		var element UserIn.UserIn
		err := data.db.ScanRows(rows, &element)
		if err != nil {
			log.Info().Msg(err.Error())
		}
		dbUsers = append(dbUsers, element)
	}

	return dbUsers, nil
}

func (data *UsersReposWorker) GetUserByUsername(username string) (User.User, error) {
	var user User.User

	if err := data.db.Table(users).Where("username = ?", username).Scan(&user).Error; err != nil {
		return User.User{}, fmt.Errorf("SQL : Cannot select from table %s: Reason : %s", users, err.Error())
	}
	if user.Username == "" {
		log.Error().Msgf("SQL : No rows in result set with username : %s", username)
	}
	return user, nil
}

func InitUsersRepos(db *gorm.DB) *UsersReposWorker {
	return &UsersReposWorker{db: *db}
}
