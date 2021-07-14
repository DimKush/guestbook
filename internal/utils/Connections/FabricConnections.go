package DbConnections

import (
	"sync"

	"github.com/DimKush/guestbook/tree/main/internal/Configurator"
	"gorm.io/gorm"
)

type Connector interface{}

var instance *multiConnector = nil
var once sync.Once

// TODO : for now connection is only one. It's a postgres
type multiConnector struct {
	connections map[string]Connection
}

func Instance() Connector {
	once.Do(func() {
		if instance == nil {
			instance = new(multiConnector)
			// TODO for now only for a postgres db
			if Configurator.Instance().GetDbCore() == "postgres" {
				instance.setPgConnector()
			}
		}
	})
	return instance
}

func (data *multiConnector) GetPgConnection() *gorm.DB {
	return data.connections["postgres"].GetDbConnection()

}

func (data *multiConnector) setPgConnector() {
	data.connections["postgres"] = newPgConnection()
}
