package service

import (
	"strings"
	"time"

	"github.com/DimKush/guestbook/tree/main/internal/entities/AuditEvent"
	"github.com/DimKush/guestbook/tree/main/pkg/repository"
)

const (
	AUDIT_FATAL = iota
	AUDIT_ERROR
	AUDIT_WARNING
	AUDIT_INFO
	AUDIT_DEBUG
	AUDIT_TRACE
)

type Audit interface {
	WriteEvent() error
}

type AuditService struct {
	Audit
	CurrentLevel int
}

func InitAudit(log_level string, repos *repository.Repository) *AuditService {
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
		Audit:        repos.Audit,
		CurrentLevel: cur_lvl,
	}
}

func (data *AuditService) WriteEvent(service_name string, initiator string, event_type int, event_date time.Time, is_panic bool, description string) {
	data.WriteEventStruct(AuditEvent.AuditEvent{
		ServiceName: service_name,
		Initiator:   initiator,
		//EventType:   event_type,
	})
}

func (data *AuditService) WriteEventStruct(event AuditEvent.AuditEvent) {

}
