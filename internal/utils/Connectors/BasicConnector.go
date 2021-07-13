package DbConnections

import "gorm.io/gorm"

type BasicConnection struct {
	DbConnector *gorm.DB
	DbType      string
}

func (data *BasicConnection) getDbConnection() *gorm.DB {
	return data.DbConnector
}
