package DbConnectors

import (
	"fmt"
	"sync"

	"github.com/DimKush/guestbook/tree/main/internal/Configurator"
	"github.com/DimKush/guestbook/tree/main/internal/Logger"
)

type ConnectionPool interface{}

var instance *connections = nil
var once sync.Once

func Instance() ConnectionPool {
	once.Do(func() {
		if instance != nil {
			instance = new(connections)
		}
	})

	return instance
}

type connections struct {
}

func SetConnection(connectionType int) (ConnectionUnits, error) {
	if connectionType < Configurator.DB_POSTGRES || connectionType > Configurator.DB_SQLITE {
		err := fmt.Errorf("Incorrect type of database connection.")
		Logger.Instance().Log().Fatal().Msg(err.Error())

		return nil, err
	}

	switch connectionType {
	case Configurator.DB_POSTGRES:
		{
			pg_connection, err := NewPgConnection()

			if err != nil {
				Logger.Instance().Log().Fatal().Msg(err.Error())
				return nil, fmt.Errorf("")
			}

			return pg_connection, nil
		}
	default:
		{
			err := fmt.Errorf("Error, Cannot create connection.")
			Logger.Instance().Log().Fatal().Msg(err.Error())

			return nil, err
		}
	}
}
