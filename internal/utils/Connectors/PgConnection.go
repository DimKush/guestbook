package DbConnections

import (
	"github.com/DimKush/guestbook/tree/main/internal/Configurator"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Connection interface {
}

type pgConnection struct {
	connector BasicConnection
}

func (data *pgConnection) open() error {
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

func newPgConnection() (pgConnection, error) {
	var connect pgConnection
	err := connect.open()

	if err != nil {
		return pgConnection{}, err
	} else {
		return connect, nil
	}
}
