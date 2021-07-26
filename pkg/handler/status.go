package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func (h *Handler) status(context *gin.Context) {
	log.Error().Msg("Status")
}
