package service

import (
	"strings"

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
	WriteEvent(event AuditEvent.AuditEvent) error
}

type AuditService struct {
	Audit
	CurrentLevel int
}

func InitAudit(repos repository.Audit, log_level string) AuditService {
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
	return AuditService{
		Audit:        AuditWriterInit(repos),
		CurrentLevel: cur_lvl,
	}
}
