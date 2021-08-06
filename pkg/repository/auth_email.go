package repository

import (
	"github.com/DimKush/guestbook/tree/main/internal/entities/EmailEventDb"
	"gorm.io/gorm"
)

type EmailEventRep struct {
	db *gorm.DB
}

func (data *EmailEventRep) InitEmailEvent(email_event EmailEventDb.EmailEventDb) error {
	return nil
}

func InitEmailEventRep(database *gorm.DB) *EmailEventRep {
	return &EmailEventRep{
		db: database,
	}
}
