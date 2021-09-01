package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/DimKush/guestbook/tree/main/internal/entities/EventItem"
	"github.com/DimKush/guestbook/tree/main/internal/entities/UserIn"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func (h *Handler) getEventsByParams(context *gin.Context) {
	log.Info().Msg("Handler getEventsByParams process request.")
	// get current user
	h.userIdentity(context)
	var user UserIn.UserIn
	if userId, exists := context.Get(userCTX); exists {
		if convertedId, ok := userId.(int); !ok {
			err := fmt.Errorf("Error parsing userId %v", userId)
			log.Error().Msgf("Error during parsing userIdentity : %s", err.Error())
			initErrorResponce(context, http.StatusBadRequest, err.Error())
		} else {
			user.Id = convertedId
		}
	} else {
		err := fmt.Errorf("Incorrect current username.")
		log.Error().Msgf("Error during parsing json : %s", err.Error())
		initErrorResponce(context, http.StatusBadRequest, err.Error())
		return
	}

	// parse json
	var event EventItem.EventItem
	if err := context.BindJSON(&event); err != nil {
		log.Error().Msgf("Error during parsing json : %s", err.Error())
		initErrorResponce(context, http.StatusBadRequest, err.Error())
		return
	}

	event.EventOwnerId = user.Id

	events, err := h.services.GetEventsByParams(event)
	if err != nil {
		initErrorResponce(context, http.StatusBadRequest, "")
		return
	}

	initOkResponce(context, map[string]interface{}{
		"Result": events,
	})
}

func (h *Handler) createEvent(context *gin.Context) {
	fmt.Println("WOW")
	log.Info().Msg("Handler createEvent process request.")

	list_id, err := strconv.Atoi(context.Param("list_id"))
	if err != nil {
		initErrorResponce(context, http.StatusBadRequest, "Incorrect list_id in the url.")
		return
	}
	if list_id == 0 {
		initErrorResponce(context, http.StatusBadRequest, "In the url list_id cannot be 0.")
		return
	}

	fmt.Println(list_id)
	initOkResponce(context, map[string]interface{}{})
}

func (h *Handler) getAllEvents(context *gin.Context) {

}

func (h *Handler) getEventById(context *gin.Context) {

}

func (h *Handler) updateEventById(context *gin.Context) {

}

func (h *Handler) dropEventById(context *gin.Context) {

}
