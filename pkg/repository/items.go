package repository

import (
	"fmt"

	"github.com/DimKush/guestbook/tree/main/internal/entities/Item"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

type ItemsRepo struct {
	db gorm.DB
}

func (data *ItemsRepo) GetItemsByParams(item Item.Item) ([]Item.Item, error) {
	var allEvents []Item.Item

	query := data.db.Table(items).Select("items.*, lst.title as list_title, tp.fullname as event_type_name").
		Joins("left join events_lists lst on lst.id = items.list_id").
		Joins("left join item_type tp on tp.type_id = items.item_type_id").
		Joins("left join users us on us.id = lst.owner_user_id")

	query.Where("us.id = ?", item.ItemOwnerId)

	if (item != Item.Item{}) {
		if item.ListId != 0 {
			query.Where("lst.id = ?", item.ListId)
		}
		if item.Id != 0 {
			query.Where("items.id = ?", item.Id)
		}
		if item.ListTitle != "" {
			listLike := "%" + item.ListTitle + "%"
			query.Where("lst.title like ?", listLike)
		}
		if item.ItemTypeName != "" {
			eventTypeLike := "%" + item.ItemTypeName + "%"
			query.Where("tp.fullname like ?", eventTypeLike)
		}
		if item.Description != "" {
			description := "%" + item.Description + "%"
			query.Where("items.description like ?", description)
		}
	}

	rows, err := query.Rows()
	if err != nil {
		log.Error().Msg(err.Error())
		return nil, fmt.Errorf("Error during execute query.")
	}

	var event Item.Item

	for rows.Next() {
		data.db.ScanRows(rows, &event)
		allEvents = append(allEvents, event)
	}

	return allEvents, nil
}

func InitItemsRep(database *gorm.DB) *ItemsRepo {
	return &ItemsRepo{
		db: *database,
	}
}