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
