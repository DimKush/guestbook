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

	query := data.db.Table(items).Select("items.*, lst.title as list_title, tp.fullname as item_type_name").
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

	rows, err := query.Debug().Rows()
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

func (data *ItemsRepo) CreateNewItem(item Item.Item) error {
	err := data.db.Table(items).Select("list_id", "item_type_id", "description").Create(&item).Error
	if err != nil {
		return fmt.Errorf("Error during execute query in database.")
	} else {
		return nil
	}
}

func (data *ItemsRepo) GetItemTypesByParams(item Item.ItemType) ([]Item.ItemType, error) {
	var item_types []Item.ItemType
	query := data.db.Table(item_type)

	if item.TypeId != 0 {
		query.Where("type_id = ?", item.TypeId)
	}
	if item.Systemname != "" {
		likeConst := "%" + item.Systemname + "%"
		query.Where("systemname like ?", likeConst)
	}
	if item.Fullname != "" {
		likeConst := "%" + item.Fullname + "%"
		query.Where("fullname like ?", likeConst)
	}

	rows, err := query.Rows()
	if err != nil {
		log.Error().Msg(err.Error())
		return nil, fmt.Errorf("Error during execute query.")
	}

	var element Item.ItemType
	for rows.Next() {
		data.db.ScanRows(rows, &element)
		item_types = append(item_types, element)
	}

	return item_types, nil
}

func (data *ItemsRepo) GetItemById(item_id int) (Item.Item, error) {
	var item Item.Item
	data.db.Table(items).Where("id = ?", item_id).Scan(&item)

	if (item == Item.Item{}) {
		return Item.Item{}, fmt.Errorf("Cannot find item with id = %d", item_id)
	}

	return item, nil
}

func (data *ItemsRepo) UpdateItemById(item *Item.Item) error {
	err := data.db.Debug().Table(items).Where("id = ?", item.Id).Updates(Item.Item{ItemTypeName: item.ItemTypeName, Description: item.Description}).Error

	return err
}

func (data *ItemsRepo) DeleteItemById(item_id int) error {
	err := data.db.Table(items).Delete(&Item.Item{}, item_id).Error

	return err
}

func (data *ItemsRepo) GetItemsAvailability(list_id int) (int64, error) {
	var items_count int64

	err := data.db.Debug().Table(items).Where("list_id = ?", list_id).Count(&items_count).Error

	if err != nil {
		return 0, err
	} else {
		return items_count, nil
	}
}

func InitItemsRep(database *gorm.DB) *ItemsRepo {
	return &ItemsRepo{
		db: *database,
	}
}
