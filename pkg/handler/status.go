package handler

import (
	"github.com/gin-gonic/gin"
)

func (h *Handler) status(context *gin.Context) {
	initOkResponce(context, map[string]interface{}{
		"Message": "Server is online.",
	})
}
