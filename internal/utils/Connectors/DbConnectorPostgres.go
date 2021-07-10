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

func (data *PgConnector) Open() error {
	dialector := postgres.New(postgres.Config{
		DSN:                  Configurator.Instance().GetDbConnectGorm(Configurator.DB_POSTGRES),
		PreferSimpleProtocol: true, // // disables implicit prepared statement usage. By default pgx automatically uses the extended protocol
	})

	var err error

	data.connector.DbConnector, err = gorm.Open(dialector, &gorm.Config{})

	if err != nil {
		return err
	} else {
		data.connector.DbConnector = new(gorm.DB)
		return nil
	}

}

func NewPgConnection() (PgConnector, error) {
	var connect PgConnector
	err := connect.Open()

	if err != nil {
		return PgConnector{}, err
	} else {
		return connect, nil
	}
}
