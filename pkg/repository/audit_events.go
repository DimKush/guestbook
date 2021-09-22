package repository

import (
	"github.com/DimKush/guestbook/tree/main/internal/entities/AuditEvent"
	"gorm.io/gorm"
)

type AuditEventRep struct {
	db *gorm.DB
}

func InitAuditRep(database *gorm.DB) *AuditEventRep {
	return &AuditEventRep{db: database}
}

func (data *AuditEventRep) WriteEvent(event AuditEvent.AuditEvent) error {
	return data.db.Table(audit_events).Create(&event).Error
}

func (data *AuditEventRep) GetAuditEventByParams(filters *AuditEvent.AuditEvent) ([]AuditEvent.AuditEvent, error) {
	var allEvents []AuditEvent.AuditEvent
	query := data.db.Table(audit_events)

	if (*filters != AuditEvent.AuditEvent{}) {
		if filters.EventId != 0 {
			query.Where("event_id = ?", filters.EventId)
		}
		if filters.ServiceName != "" {

			query.Where("service_name like ?", string("%"+filters.ServiceName+"%"))
		}
		if filters.Initiator != "" {
			query.Where("initiator like ?", string("%"+filters.Initiator+"%"))
		}
		if filters.EventType != "" {
			query.Where("event_type like ?", string("%"+filters.EventType+"%"))
		}
		// if filters.EventDate != tim {}
		if filters.Description != "" {
			query.Where("description like ?", string("%"+filters.Description+"%"))
		}
	}

	rows, err := query.Rows()
	if err != nil {
		return nil, err
	}

	var audit AuditEvent.AuditEvent
	for rows.Next() {
		data.db.ScanRows(rows, &audit)
		allEvents = append(allEvents, audit)
	}

	return allEvents, nil
}
