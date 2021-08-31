package service

import (
	"sort"

	"github.com/DimKush/guestbook/tree/main/internal/entities/EventItem"
	"github.com/DimKush/guestbook/tree/main/pkg/repository"
)

type EventsServiceWorker struct {
	events_repo repository.EventsService
}

func (data *EventsServiceWorker) GetEventsByParams(event EventItem.EventItem) ([]EventItem.EventItem, error) {
	events, err := data.events_repo.GetEventsByParams(event)

	if err != nil {
		return nil, err
	}

	if len(events) != 0 {
		sort.Slice(events, func(i int, j int) bool {
			return events[i].Id < events[j].Id
		})
	} else {
		events = []EventItem.EventItem{}
	}

	return events, nil
}

func InitEventsServiceWorker(events repository.EventsService) *EventsServiceWorker {
	return &EventsServiceWorker{
		events_repo: events,
	}
}
