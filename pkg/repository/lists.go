package repository

import (
	"github.com/DimKush/guestbook/tree/main/internal/entities/List"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

type ListServiceRepo struct {
	db *gorm.DB
}

func (data *ListServiceRepo) GetAllLists() ([]List.List, error) {

	var allLists []List.List

	rows, err := data.db.Table(events_lists).Select("events_lists.*, users.username as owner").Joins("left join users on users.id=owner_user_id").Rows()
	defer rows.Close()

	if err != nil {
		return []List.List{}, err
	}

	for rows.Next() {
		var element List.List
		data.db.ScanRows(rows, &element)
		allLists = append(allLists, element)
	}

	return allLists, nil
}

func (data *ListServiceRepo) GetListsByParams(list List.List) ([]List.List, error) {
	log.Info().Msgf("%v\n", list)
	var allLists []List.List

	query := data.db.Table(events_lists).Select("events_lists.*, users.username as owner").Joins("left join users on users.id=owner_user_id")

	if (list != List.List{}) {
		if list.Id != 0 {
			query.Where("events_lists.id=?", list.Id)
		}
		if list.Title != "" {
			title := "%" + list.Title + "%"
			query.Where("events_lists.title LIKE ?", title)
		}
		if list.Description != "" {
			description := "%" + list.Description + "%"
			query.Where("events_lists.description LIKE ?", description)
		}
		if list.Owner != "" {
			owner := "%" + list.Owner + "%"
			query.Where("users.username LIKE ?", owner)
		}
	}

	rows, err := query.Rows()

	defer rows.Close()

	if err != nil {
		return []List.List{}, err
	}

	for rows.Next() {
		var element List.List
		data.db.ScanRows(rows, &element)
		allLists = append(allLists, element)
	}

	return allLists, nil
}

func InitListsRep(database *gorm.DB) *ListServiceRepo {
	return &ListServiceRepo{db: database}
}
