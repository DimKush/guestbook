package AuditEvent

import "time"

type AuditEvent struct {
	Id          int       `gorm:"event_id"`
	ServiceName string    `gorm:"description"`
	Initiator   string    `gorm:"initiator"`
	EventType   string    `gorm:"service_name"`
	EventDate   time.Time `gorm:"event_date"`
	IsPanic     bool      `gorm:"is_panic"`
	Description string    `gorm:"event_type"`
}
