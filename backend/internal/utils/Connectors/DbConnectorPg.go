package connectors

import (
	"github.com/DimKush/guestbook/tree/main/backend/internal/Configurator"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PgConnector struct {
	pgdb *gorm.DB // database connector
}

func (data *PgConnector) Open() {
	dialector := postgres.New(postgres.Config{
		DSN:                  Configurator.Instance().GetDbConnectGorm(Configurator.DB_POSTGRES),
		PreferSimpleProtocol: true, // // disables implicit prepared statement usage. By default pgx automatically uses the extended protocol
	})

	var err error

	data.pgdb, err = gorm.Open(dialector, &gorm.Config{})
	
	
	if err != nil {
		// TODO when will be another databases turn off panic and try to connect to another db
		//panic()
	} else {
		data.pgdb = new(gorm.DB)
	}

	//check if all alright
	data.pgdb.DB()
}
