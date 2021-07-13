package DbConnections

import "gorm.io/gorm"

var instance *Connection = nil

type MultiConnection struct {
	connections map[string]*Connection
}

func (data *MultiConnection) GetPgConnection() *gorm.DB {

	data.connections["postgres"] = newPgConnection()
}
