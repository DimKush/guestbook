package repository

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	users        = "users"
	events_lists = "events_lists"
	users_lists  = "users_lists"
	event_item   = "event_item"
	audit_events = "audit_events"
)

var system_tables = initSystemTablesConf(db_type_postgres)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	Dbname   string
	Timezone string
	SSLMode  string
}

func NewPostgresConnection(cfg Config) (*gorm.DB, error) {
	var DSN string = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		cfg.Host,
		cfg.Username,
		cfg.Password,
		cfg.Dbname,
		cfg.Port,
		cfg.SSLMode,
		cfg.Timezone)

	fmt.Printf("\nDSN : %s", DSN)
	dialector := postgres.New(postgres.Config{
		DSN:                  DSN,
		PreferSimpleProtocol: true,
	})

	return gorm.Open(dialector, &gorm.Config{})
}
