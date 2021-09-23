package handler

import (
	"net/http"

	"github.com/DimKush/guestbook/tree/main/internal/entities/User"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func (h *Handler) getAllUsernames(context *gin.Context) {
	log.Info().Msg("Handler GetAllUsernames process request.")

	users, err := h.services.GetAllUsernames()

	if err != nil {
		initErrorResponce(context, http.StatusInternalServerError, err.Error())
	}

	initOkResponce(context, map[string]interface{}{
		"Result": users,
	})
}

func (h *Handler) getUsersByParams(context *gin.Context) {
	log.Info().Msg("Handler getUsersByParams process request.")

	var user User.User

	if err := context.BindJSON(&user); err != nil {
		initErrorResponce(context, http.StatusInternalServerError, err.Error())
		return
	}

	_, err := h.services.GetUsersByParams(&user) // TODO : error

	if err != nil {
		initErrorResponce(context, http.StatusInternalServerError, err.Error())
		return
	}

}
