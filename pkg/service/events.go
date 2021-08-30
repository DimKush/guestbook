package service

import "github.com/DimKush/guestbook/tree/main/pkg/repository"

type EventsService struct {
	events_repo repository.EventsRepo
}
