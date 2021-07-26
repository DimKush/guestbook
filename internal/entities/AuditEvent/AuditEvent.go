package AuditEvent

import "time"

type AuditEvent struct {
	Id               int       `gorm:"event_id"`
	Description      string    `gorm:"event_type"`
	EventDate        time.Time `gorm:"event_date"`
	EventType        string    `gorm:"service_name"`
	OperationSysName string    `gorm:"initiator"`
	IsPanic          bool      `gorm:"IsPanic"`
	ServiceName      string    `gorm:"description"`
}
