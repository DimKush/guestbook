package repository

import (
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

func InitUsersRepos(db *gorm.DB) *UsersReposWorker {
	return &UsersReposWorker{db: *db}
}
