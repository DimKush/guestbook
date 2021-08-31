package repository

import (
	"github.com/DimKush/guestbook/tree/main/internal/entities/EventItem"
	"gorm.io/gorm"
)

type EventsRepo struct {
	db gorm.DB
}

func (data *EventsRepo) GetEventsByParams(item EventItem.EventItem) ([]EventItem.EventItem, error) {
	query := data.db.Debug().Table(event_item).Select("event_item.*, event_type.type_id as fullname, events_lists.title").
		Joins("left join events_lists lst on lst.id = event_item.list_id").
		Joins("left join event_type tp on tp.id = event_item.event_type_id").
		Joins("left hoin users us on us.id = lst.owner_user_id")

	query.Where("us.id = ?", item.EventOwnerId)

	if (item != EventItem.EventItem{}) {
		if item.Id != 0 {
			query.Where("event_item.id = ?", item.Id)
		}
		if item.ListTile != "" {
			listLike := "%" + item.ListTile + "%"
			query.Where("lst.title like ?", listLike)
		}
		if item.EventTypeName != "" {
			eventTypeLike := "%" + item.EventTypeName + "%"
			query.Where("tp.fullname like ?", eventTypeLike)
		}
		if item.Description != "" {
			description := "%" + item.EventTypeName + "%"
			query.Where("event_item.description like ?", description)
		}
	}

	return nil, nil
}

func InitEventsRep(database *gorm.DB) *EventsRepo {
	return &EventsRepo{
		db: *database,
	}
}
