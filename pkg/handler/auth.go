package handler

import (
	"net/http"

	"github.com/DimKush/guestbook/tree/main/internal/entities/User"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func (h *Handler) signUp(context *gin.Context) {
	var user User.User
	log.Info().Msg("signIn process request.")
	if err := context.BindJSON(&user); err != nil {
		initErrorResponce(context, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.services.Authorization.CreateUser(user)
	if err != nil {
		initErrorResponce(context, http.StatusInternalServerError, err.Error())
		return
	}

	initOkResponce(context, map[string]interface{}{
		"id": id,
	})

}

type signInInput struct {
}

func (h *Handler) signIn(context *gin.Context) {

}
