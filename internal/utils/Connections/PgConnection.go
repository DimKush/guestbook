package DbConnections

import (
	"github.com/DimKush/guestbook/tree/main/internal/Configurator"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type pgConnection struct {
	connector BasicConnection
}

func (data *pgConnection) open() error {
	dialector := postgres.New(postgres.Config{
		DSN:                  Configurator.Instance().GetDbConnectGorm(Configurator.DB_POSTGRES),
		PreferSimpleProtocol: true, // // disables implicit prepared statement usage. By default pgx automatically uses the extended protocol
	})

	var err error
	data.connector.DbType = "postgres"
	data.connector.DbConnector, err = gorm.Open(dialector, &gorm.Config{})

	if err != nil {
		return err
	} else {
		data.connector.DbConnector = new(gorm.DB)
		return nil
	}
}

func (data *pgConnection) GetDbConnection() *gorm.DB {
	return data.GetDbConnection()
}

func newPgConnection() Connection {
	var connect pgConnection
	err := connect.open()

	if err != nil {
		return &connect
	} else {
		return &BasicConnection{}
	}

}
