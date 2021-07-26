package repository

import "gorm.io/gorm"

type AuditEventRep struct {
	db *gorm.DB
}

func InitAuditRep(database *gorm.DB) *AuditEventRep {
	return &AuditEventRep{db: database}
}

func (data *AuditEventRep) WriteEvent() error {
	return nil
}

func (data *AuditEventRep) WriteEventStruct() error {
	return nil
}
