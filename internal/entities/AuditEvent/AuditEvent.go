package AuditEvent

import "time"

type AuditEvent struct {
	ServiceName string    `gorm:"service_name"`
	Initiator   string    `gorm:"initiator"`
	EventType   string    `gorm:"event_type"`
	EventDate   time.Time `gorm:"event_date"`
	IsPanic     bool      `gorm:"is_panic"`
	Description string    `gorm:"description"`
}
