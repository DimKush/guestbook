package handler

import (
	"net/http"

	"github.com/DimKush/guestbook/tree/main/internal/entities/User"
	"github.com/DimKush/guestbook/tree/main/internal/entities/UserIn"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func (h *Handler) signUp(context *gin.Context) {
	var user User.User
	log.Info().Msg("signUp process request.")
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

func (h *Handler) signIn(context *gin.Context) {
	var userIn UserIn.UserIn
	log.Info().Msg("signIn process request.")

	if err := context.BindJSON(&userIn); err != nil {
		initErrorResponce(context, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.services.Authorization.GenerateToken(userIn.Username, userIn.Password)
	if err != nil {
		log.Error().Msg(err.Error())
		initErrorResponce(context, http.StatusBadRequest, err.Error())
		return
	}

	initOkResponce(context, map[string]interface{}{
		"token": token,
	})
}
