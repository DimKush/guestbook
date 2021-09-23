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

func (data *UsersReposWorker) GetUsersByParams(filter *User.User) ([]User.User, error) {
	var usersArr []User.User

	query := data.db.Table(users)
	if (*filter != User.User{}) {
		if filter.Id != 0 {
			query.Where("id = ?", filter.Id)
		}
		if filter.Name != "" {
			query.Where("name like ?", string("%"+filter.Name+"%"))
		}
		if filter.Username != "" {
			query.Where("username like ?", string("%"+filter.Username+"%"))
		}
		if filter.Email != "" {
			query.Where("email like ?", string("%"+filter.Email+"%s"))
		}
		// if filter.Registration_date TODO registration date
	}

	rows, err := query.Rows()
	if err != nil {
		return nil, err
	}

	var element User.User
	for rows.Next() {
		data.db.ScanRows(rows, &element)
		usersArr = append(usersArr, element)
	}

	return usersArr, nil
}

func InitUsersRepos(db *gorm.DB) *UsersReposWorker {
	return &UsersReposWorker{db: *db}
}
