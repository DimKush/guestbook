package repository

import (
	"fmt"

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

func (data *ListServiceRepo) GetAutoListId() (int, error) {
	var idVal int

	data.db.Table(system_tables.sequences).Select("last_value as id").Where("sequencename = ?", "events_lists_id_seq").Scan(&idVal)
	if idVal == 0 {
		return 0, fmt.Errorf("Can't get nextval('events_lists_id_seq')")
	}

	return idVal, nil
}

func (data *ListServiceRepo) GetListById(list_id int) (List.List, error) {
	var list List.List

	data.db.Table(events_lists).Select("events_lists.*, users.username as owner").
		Joins("left join users on users.id=owner_user_id").
		Where("events_lists.id=?", list_id).
		Scan(&list)

	if (list == List.List{}) {
		return List.List{}, fmt.Errorf("Cannot find list with id = %d", list_id)
	}

	return list, nil
}

func (data *ListServiceRepo) CreateList(list List.List) error {
	err := data.db.Table(events_lists).Create(&list).Error
	if err != nil {
		log.Error().Msgf(err.Error())
		return err
	}

	return nil
}

func InitListsRep(database *gorm.DB) *ListServiceRepo {
	return &ListServiceRepo{db: database}
}
