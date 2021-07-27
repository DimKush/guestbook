


type struct AuditWriter{
	audit repository.Audit
}


func AuditWriterInit(repos *repository.Audit) service.Audit{
	return &AuditWriter{
		audit : repos.Audit
	}
}

func (data *AuditWriter) WriteEvent(event AuditEvent.AuditEvent) error {
	return nil
	//return Audit.WriteEvent(event)
}
