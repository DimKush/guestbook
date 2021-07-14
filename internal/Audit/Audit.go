package Audit

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/DimKush/guestbook/tree/main/internal/Configurator"
	"github.com/DimKush/guestbook/tree/main/internal/Logger"
)

const (
	AUDIT_FATAL   = iota
	AUDIT_ERROR   = iota
	AUDIT_WARNING = iota
	AUDIT_INFO    = iota
	AUDIT_DEBUG   = iota
	AUDIT_TRACE   = iota
)

type Controller interface {
	Execute(writer http.ResponseWriter, reader *http.Request)
}

type Audit struct {
	EventType   string    `json:"EventType"`
	EventDate   time.Time `json:"EventDate"`
	ServiceName string    `json:"ServiceName"`
	IsPanic     bool      `json:"IsPanic"`
	Description string    `json:"Description"`
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

	fmt.Printf("%v", data)
}

func (data *Audit) returnEventType(typeStr string) int {
	config_audit_level := strings.ToLower(typeStr)

	switch config_audit_level {
	case "fatal":
		{
			return AUDIT_FATAL
		}
	case "error":
		{
			return AUDIT_ERROR
		}
	case "warning":
		{
			return AUDIT_WARNING
		}
	case "info":
		{
			return AUDIT_INFO
		}
	case "debug":
		{
			return AUDIT_DEBUG
		}
	case "trace":
		{
			return AUDIT_TRACE
		}
	default:
		{
			return AUDIT_ERROR
		}
	}
}

func (data *Audit) ExecuteQuery() {
	DbConnections.Instance()
}

func NewAudit() Controller {
	return &Audit{}
}
