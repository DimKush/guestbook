package handler

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/DimKush/guestbook/tree/main/internal/entities/UserIn"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

const (
	authorizationHandler = "Authorization"
	userCTX              = "userId"
)

func (h *Handler) userIdentity(context *gin.Context) {
	header := context.GetHeader(authorizationHandler)

	if header == "" {
		initErrorResponce(context, http.StatusUnauthorized, fmt.Sprint("Empty auth header."))
		return
	}

	headerAuthParts := strings.Split(header, " ")
	if len(headerAuthParts) != 2 {
		initErrorResponce(context, http.StatusUnauthorized, fmt.Sprint("Bad header. Cannot format."))
		return
	}

	userId, err := h.services.ParseToken(headerAuthParts[1])
	if err != nil {
		log.Error().Msgf("Error during userIdentity. Reason : %s", err.Error())
		initErrorResponce(context, http.StatusUnauthorized, err.Error())
		return
	}

	context.Set(userCTX, userId)
}

func (h *Handler) userIdentityToken(context *gin.Context) {
	cookie_token, err := context.Cookie("jwt")
	if err != nil {
		log.Error().Msg(err.Error())
	}

	userId, err := h.services.ParseToken(cookie_token)
	if err != nil {
		log.Error().Msgf("Error during userIdentity. Reason : %s", err.Error())
		initErrorResponce(context, http.StatusUnauthorized, err.Error())
		return
	}

	user, err := h.services.GetUser(UserIn.UserIn{Id: userId})

	if err != nil {
		initErrorResponce(context, http.StatusUnauthorized, err.Error())
	}

	initOkResponce(context, map[string]interface{}{
		"Status":   "OK",
		"Username": user.Username,
	})
}

func (h *Handler) logout(context *gin.Context) {
	initOkResponce(context, map[string]interface{}{})
}
