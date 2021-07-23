package AuditEvent

import "time"

type AuditEvent struct {
	Description      string    `json:"Description"`
	EventDate        time.Time `json:"EventDate"`
	EventType        string    `json:"EventType"`
	IsPanic          bool      `json:"IsPanic"`
	OperationSysName string    `json:"OperationSysName"`
	ServiceName      string    `json:"ServiceName"`
}
