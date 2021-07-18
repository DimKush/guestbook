package Audit

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/DimKush/guestbook/tree/main/internal/AuditProxy"
	"github.com/DimKush/guestbook/tree/main/internal/Configurator"
	"github.com/DimKush/guestbook/tree/main/internal/Logger"
	DbConnections "github.com/DimKush/guestbook/tree/main/internal/utils/Connections"
)

type Controller interface {
	Execute(writer http.ResponseWriter, reader *http.Request)
}

type Audit struct {
	EventType   string    `gorm:"column:eventtype"`
	EventDate   time.Time `gorm:"column:eventdate"`
	ServiceName string    `gorm:"column:servicename"`
	IsPanic     bool      `gorm:"column:is_panic"`
	Description string    `gorm:"column:description"`
}

func (data *Audit) Execute(writer http.ResponseWriter, reader *http.Request) {
	Logger.Instance().Log().Info().Msg("AuditEvent Execute process request")

	err := json.NewDecoder(reader.Body).Decode(&data)
	if err != nil {
		Logger.Instance().Log().Error().Msgf("Cannot parse json %s", err.Error())
	}

	currentEventType := data.returnEventType(data.EventType)
	confEventType := data.returnEventType(Configurator.Instance().GetAuditLevel())

	if currentEventType > confEventType {
		fmt.Printf("%d > %d", currentEventType, confEventType)
		return
	}

	//TODO : for test
	data.ExecuteQuery()
}

func (data *Audit) returnEventType(typeStr string) int {
	config_audit_level := strings.ToLower(typeStr)

	switch config_audit_level {
	case "fatal":
		{
			return AuditProxy.AUDIT_FATAL
		}
	case "error":
		{
			return AuditProxy.AUDIT_ERROR
		}
	case "warning":
		{
			return AuditProxy.AUDIT_WARNING
		}
	case "info":
		{
			return AuditProxy.AUDIT_INFO
		}
	case "debug":
		{
			return AuditProxy.AUDIT_DEBUG
		}
	case "trace":
		{
			return AuditProxy.AUDIT_TRACE
		}
	default:
		{
			return AuditProxy.AUDIT_ERROR
		}
	}
}

func (data *Audit) ExecuteQuery() {
	//TODO :  assignment to entry in nil map
	connecter := DbConnections.Instance().GetPgConnection()
	connecter.Table("public.audit_events").Create(&data)
}

func NewAudit() Controller {
	return &Audit{}
}
