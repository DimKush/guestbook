package repository

import (
	"fmt"

	"github.com/DimKush/guestbook/tree/main/internal/entities/EventItem"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

type EventsRepo struct {
	db gorm.DB
}

func (data *EventsRepo) GetEventsByParams(item EventItem.EventItem) ([]EventItem.EventItem, error) {
	var allEvents []EventItem.EventItem

	query := data.db.Table(event_item).Select("event_item.*, lst.title as list_title, tp.fullname as event_type_name").
		Joins("left join events_lists lst on lst.id = event_item.list_id").
		Joins("left join event_type tp on tp.type_id = event_item.event_type_id").
		Joins("left join users us on us.id = lst.owner_user_id")

	query.Where("us.id = ?", item.EventOwnerId)

	if (item != EventItem.EventItem{}) {
		if item.Id != 0 {
			query.Where("event_item.id = ?", item.Id)
		}
		if item.ListTitle != "" {
			listLike := "%" + item.ListTitle + "%"
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

	rows, err := query.Rows()
	if err != nil {
		log.Error().Msg(err.Error())
		return nil, fmt.Errorf("Error during execute query.")
	}

	var event EventItem.EventItem

	for rows.Next() {
		data.db.ScanRows(rows, &event)
		allEvents = append(allEvents, event)
	}

	return allEvents, nil
}

func InitEventsRep(database *gorm.DB) *EventsRepo {
	return &EventsRepo{
		db: *database,
	}
}
