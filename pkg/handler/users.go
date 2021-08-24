package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func (h *Handler) getAllUsernames(context *gin.Context) {
	fmt.Println("LOL")
	log.Info().Msg("Handler GetAllUsernames process request.")

	users, err := h.services.GetAllUsernames()

	if err != nil {
		initErrorResponce(context, http.StatusInternalServerError, err.Error())
	}

	initOkResponce(context, map[string]interface{}{
		"Result": users,
	})
}
