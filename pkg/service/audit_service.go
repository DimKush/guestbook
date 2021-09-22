package service

import (
	"fmt"
	"strings"
	"time"

	"github.com/DimKush/guestbook/tree/main/internal/entities/AuditEvent"
	"github.com/DimKush/guestbook/tree/main/pkg/repository"
	"github.com/rs/zerolog/log"
)

const (
	AUDIT_FATAL = iota
	AUDIT_ERROR
	AUDIT_WARNING
	AUDIT_INFO
	AUDIT_DEBUG
	AUDIT_TRACE
)

type AuditService struct {
	audit        repository.AuditInt
	currentLevel int
}

func InitAudit(repos *repository.Repository, log_level string) *AuditService {
	log_level = strings.ToLower(log_level)

	var cur_lvl int
	switch log_level {
	case "fatal":
		{
			cur_lvl = AUDIT_FATAL
		}
	case "error":
		{
			cur_lvl = AUDIT_ERROR
		}
	case "warning":
		{
			cur_lvl = AUDIT_WARNING
		}
	case "info":
		{
			cur_lvl = AUDIT_INFO
		}
	case "debug":
		{
			cur_lvl = AUDIT_DEBUG
		}
	case "trace":
		{
			cur_lvl = AUDIT_TRACE
		}
	default:
		{
			cur_lvl = AUDIT_ERROR
		}
	}
	return &AuditService{
		audit:        repos.AuditInt,
		currentLevel: cur_lvl,
	}
}

var Audit *AuditService = nil

func (data *AuditService) WriteEventParams(service_name string, initiator string, event_type int, event_date time.Time, is_panic bool, description string) error {
	if data.currentLevel < event_type {
		err := fmt.Errorf("Couldn't write event in a audit because current audit level is lower than input : current : %d input : %d", data.currentLevel, event_type)
		log.Error().Msgf(err.Error())

		return err
	}

	var currentLvlStr string

	switch event_type {
	case AUDIT_FATAL:
		{
			currentLvlStr = "fatal"
		}
	case AUDIT_ERROR:
		{
			currentLvlStr = "error"
		}
	case AUDIT_WARNING:
		{
			currentLvlStr = "warning"
		}
	case AUDIT_INFO:
		{
			currentLvlStr = "info"
		}
	case AUDIT_DEBUG:
		{
			currentLvlStr = "debug"
		}
	case AUDIT_TRACE:
		{
			currentLvlStr = "trace"
		}
	default:
		{
			err := fmt.Errorf("Cannot set audit level from the event_type = %d", event_type)
			log.Error().Msg(err.Error())
			return err
		}
	}

	return data.writeEvent(AuditEvent.AuditEvent{
		ServiceName: service_name,
		Initiator:   initiator,
		EventType:   currentLvlStr,
		EventDate:   event_date,
		IsPanic:     is_panic,
		Description: description,
	})
}

func (data *AuditService) WriteEventParamsBody(service_name string, initiator string, event_type int, event_date time.Time, is_panic bool, description string, body interface{}) error {
	if data.currentLevel < event_type {
		err := fmt.Errorf("Couldn't write event in a audit because current audit level is lower than input : current : %d input : %d", data.currentLevel, event_type)
		log.Error().Msgf(err.Error())

		return err
	}

	var currentLvlStr string

	switch event_type {
	case AUDIT_FATAL:
		{
			currentLvlStr = "fatal"
		}
	case AUDIT_ERROR:
		{
			currentLvlStr = "error"
		}
	case AUDIT_WARNING:
		{
			currentLvlStr = "warning"
		}
	case AUDIT_INFO:
		{
			currentLvlStr = "info"
		}
	case AUDIT_DEBUG:
		{
			currentLvlStr = "debug"
		}
	case AUDIT_TRACE:
		{
			currentLvlStr = "trace"
		}
	default:
		{
			err := fmt.Errorf("Cannot set audit level from the event_type = %d", event_type)
			log.Error().Msg(err.Error())
			return err
		}
	}

	descrWithBody := description + fmt.Sprintf(" with a body : %v", body)

	return data.writeEvent(AuditEvent.AuditEvent{
		ServiceName: service_name,
		Initiator:   initiator,
		EventType:   currentLvlStr,
		EventDate:   event_date,
		IsPanic:     is_panic,
		Description: descrWithBody,
	})
}

func (data *AuditService) writeEvent(event AuditEvent.AuditEvent) error {
	if err := data.audit.WriteEvent(event); err != nil {
		log.Error().Msgf("Cannot write %v in database. Reason : %s", event, err.Error())
		return err
	}
	return nil
}

func (data *AuditService) GetAuditEventByParams(filters *AuditEvent.AuditEvent) ([]AuditEvent.AuditEvent, error) {
	events, err := data.audit.GetAuditEventByParams(filters)
	if err != nil {
		log.Error().Msg(err.Error())
		return nil, fmt.Errorf("Error while executing in the database.")
	}

	return events, nil
}
