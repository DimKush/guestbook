package repository

import (
	"github.com/DimKush/guestbook/tree/main/internal/entities/EmailEventDb"
	"gorm.io/gorm"
)

type EmailEvent struct {
	db *gorm.DB
}

func (data *EmailEvent) CreateEmailEvent(email_event EmailEventDb.EmailEventDb) error {

	return nil
}
