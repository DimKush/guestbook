package connectors

import (
	"github.com/DimKush/guestbook/tree/main/backend/internal/Configurator"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PgConnector struct {
}

func (data *PgConnector) Open() {
	dialector := postgres.New(postgres.Config{
		DSN:                  Configurator.Instance().GetDbConnectGorm(Configurator.DB_POSTGRES),
		PreferSimpleProtocol: true, // // disables implicit prepared statement usage. By default pgx automatically uses the extended protocol
	})

	db, err := gorm.Open(dialector, &gorm.Config{})
}
