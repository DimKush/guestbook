package DbConnectors

import (
	"github.com/DimKush/guestbook/tree/main/internal/Configurator"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type ConnectionUnits interface{}

type PgConnector struct {
	connector BasicConnector
}

func (data *PgConnector) Open() {
	dialector := postgres.New(postgres.Config{
		DSN:                  Configurator.Instance().GetDbConnectGorm(Configurator.DB_POSTGRES),
		PreferSimpleProtocol: true, // // disables implicit prepared statement usage. By default pgx automatically uses the extended protocol
	})

	var err error

	data.connector.DbConnector, err = gorm.Open(dialector, &gorm.Config{})

	if err != nil {
		// TODO when will be another databases turn off panic and try to connect to another db
		//panic()
	} else {
		data.connector.DbConnector = new(gorm.DB)
	}
}

func (data *PgConnector) NewPgConnection() {

}
