package repository

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

func ServiceInit() *Repository {
	return &Repository{}
}
