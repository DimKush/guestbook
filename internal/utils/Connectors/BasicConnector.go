package DbConnectors

import "gorm.io/gorm"

type BasicConnector struct {
	DbConnector *gorm.DB
	DbType      string
}
