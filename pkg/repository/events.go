package repository

import (
	"github.com/DimKush/guestbook/tree/main/internal/entities/EventItem"
	"gorm.io/gorm"
)

type EventsRepo struct {
	db gorm.DB
}

func (data *EventsRepo) GetEventsByParams(item EventItem.EventItem) ([]EventItem.EventItem, error) {

	return nil, nil
}

func InitEventsRep(database *gorm.DB) *EventsRepo {
	return &EventsRepo{
		db: *database,
	}
}
