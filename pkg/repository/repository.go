package repository

import "gorm.io/gorm"

type Authorization interface {
}

type Event interface {
}

type EventList interface {
}

type Repository struct {
	Authorization
	Event
	EventList
}

func RepositoryInit(db *gorm.DB) *Repository {
	return &Repository{}
}
