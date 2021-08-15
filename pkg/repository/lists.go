package repository

import (
	"github.com/DimKush/guestbook/tree/main/internal/entities/List"
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
		data.db.ScanRows(rows, element)
		allLists = append(allLists, element)
	}

	return allLists, nil
}

func InitListsRep(database *gorm.DB) *ListServiceRepo {
	return &ListServiceRepo{db: database}
}
