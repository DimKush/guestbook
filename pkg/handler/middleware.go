package handler

import (
	"fmt"
	"net/http"
	"strings"

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

	context.Set("userId", userId)
}
