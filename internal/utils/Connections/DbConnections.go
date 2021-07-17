package DbConnections

import (
	"sync"

	"github.com/DimKush/guestbook/tree/main/internal/Configurator"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Connector interface {
	GetPgConnection() *gorm.DB
}

var instance *dbConnector = nil
var once sync.Once

// TODO : for now connection is only one. It's a postgres
type dbConnector struct {
	connections map[string]*gorm.DB
}

func Instance() Connector {
	once.Do(func() {
		if instance == nil {
			instance = new(dbConnector)
			if instance.connections == nil {
				instance.connections = make(map[string]*gorm.DB)
				if Configurator.Instance().GetDbCore() == "postgres" {
					instance.openPgConnection()
				}
			}
			// TODO for now only for a postgres db
		}
	})
	return instance
}

func (data *dbConnector) openPgConnection() {
	dialector := postgres.New(postgres.Config{
		DSN:                  Configurator.Instance().GetDbConnectGorm(Configurator.DB_POSTGRES),
		PreferSimpleProtocol: true, // // disables implicit prepared statement usage. By default pgx automatically uses the extended protocol
	})

	var err error
	dbConnect, err := gorm.Open(dialector, &gorm.Config{})

	if err != nil {
	} else {
		data.connections["postgres"] = dbConnect
	}
}

func (data *dbConnector) GetPgConnection() *gorm.DB {
	return data.connections["postgres"]
}
