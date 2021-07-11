package AuditEvent

import (
	"github.com/DimKush/guestbook/tree/main/internal/Logger"
)

type AuditFacade struct {
	EventType   string
	EventDate   string
	ServiceName string
	IsPanic     bool
	Description string
}

func writeAuditEvent(event AuditFacade) error {
	//	var auditRequest http.Request
	return nil
}

func writeEvent(eventType string, eventDate string, serviceName string, isPanic bool, description string) error {
	event := AuditFacade{
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
