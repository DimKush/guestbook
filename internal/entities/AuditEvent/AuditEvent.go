package AuditEvent

import "time"

type AuditEvent struct {
	EventId     int64     `json:"event_id" gorm:"event_id"`
	ServiceName string    `json:"service_name" gorm:"service_name"`
	Initiator   string    `json:"initiator" gorm:"initiator"`
	EventType   string    `json:"event_type" gorm:"event_type"`
	EventDate   time.Time `json:"event_date" gorm:"event_date"`
	IsPanic     bool      `json:"is_panic" gorm:"is_panic"`
	Description string    `json:"description" gorm:"description"`
}
