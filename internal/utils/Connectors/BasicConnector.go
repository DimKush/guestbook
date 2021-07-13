package DbConnections

import "gorm.io/gorm"

type Connection interface {
	GetDbConnection() *gorm.DB
}

type BasicConnection struct {
	DbConnector *gorm.DB
	DbType      string
}

func (data *BasicConnection) GetDbConnection() *gorm.DB {
	return data.DbConnector
}
