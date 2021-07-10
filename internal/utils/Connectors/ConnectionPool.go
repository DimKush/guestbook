package DbConnectors

import (
	"strings"
	"sync"
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

func SetConnection(connectionStr string) (ConnectionUnits, error) {
	connectStrLow := strings.ToLower(connectionStr)

	switch connectStrLow {
	case "postgres":
		{

		}
	}
}
