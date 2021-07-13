package AuditProxy

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"

	"github.com/DimKush/guestbook/tree/main/internal/Logger"
)

type AuditProxy struct {
	EventType   string
	EventDate   time.Time
	ServiceName string
	IsPanic     bool
	Description string
}

func writeAuditEvent(event AuditProxy) error {
	requestBody, err := json.Marshal(map[string]interface{}{
		"EventType":   event.EventType,
		"EventDate":   event.EventDate,
		"ServiceName": event.ServiceName,
		"IsPanic":     event.IsPanic,
		"Description": event.Description,
	})
	if err != nil {
		Logger.Instance().Log().Error().Msgf("Cannot Marshall json :%s", err.Error())
	}

	// TODO : need a port from config
	postStr := "http://localhost:5003" + "/audit/NewEvent"
	http.Post(postStr, "application/json", bytes.NewBuffer(requestBody))

	return nil
}

func WriteEvent(eventType string, eventDate time.Time, serviceName string, isPanic bool, description string) error {
	event := AuditProxy{
		EventType:   eventType,
		EventDate:   eventDate,
		ServiceName: serviceName,
		IsPanic:     isPanic,
		Description: description,
	}

	err := writeAuditEvent(event)
	if err != nil {
		Logger.Instance().Log().Error().Msgf("Error during write in audit %s", err.Error())
		return err
	}

	return nil

}
