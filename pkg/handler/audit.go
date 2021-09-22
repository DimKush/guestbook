package handler

import (
	"net/http"

	"github.com/DimKush/guestbook/tree/main/internal/entities/AuditEvent"
	"github.com/DimKush/guestbook/tree/main/pkg/service"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func (h *Handler) getAllAuditEvents(context *gin.Context) {
	log.Info().Msg("Handler getAllAuditEvents process request.")

	events, err := service.Audit.GetAuditEventByParams(&AuditEvent.AuditEvent{})
	if err != nil {
		initErrorResponce(context, http.StatusBadRequest, err.Error())
		return
	} else {
		if events != nil {
			initOkResponce(context, map[string]interface{}{
				"Result": events,
			})
		} else {
			initOkResponce(context, map[string]interface{}{})
		}
		return
	}
}

func (h *Handler) getAuditEventsByParams(context *gin.Context) {
	log.Info().Msg("Handler getAuditEventsByParams process request.")

	var filter AuditEvent.AuditEvent
	if err := context.BindJSON(&filter); err != nil {
		initErrorResponce(context, http.StatusBadRequest, err.Error())
		return
	}

	events, err := service.Audit.GetAuditEventByParams(&filter)
	if err != nil {
		initErrorResponce(context, http.StatusBadRequest, err.Error())
		return
	} else {
		if events != nil {
			initOkResponce(context, map[string]interface{}{
				"Result": events,
			})
		} else {
			initOkResponce(context, map[string]interface{}{})
		}
		return
	}

}
