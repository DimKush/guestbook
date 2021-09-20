package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/DimKush/guestbook/tree/main/internal/entities/List"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func (h *Handler) ControlListExist(context *gin.Context) (int, int, error) {
	list_id, err := strconv.Atoi(context.Param("list_id"))
	if err != nil {
		return 0, http.StatusBadRequest, fmt.Errorf("Incorrect list_id in the url.")
	}
	if list_id == 0 {
		return 0, http.StatusBadRequest, fmt.Errorf("In the url list_id cannot be 0.")
	}

	// check if list exists
	list, err := h.services.GetListById(list_id)

	if err != nil {
		return 0, http.StatusInternalServerError, fmt.Errorf("Error, during get list by id.")
	}

	if (list == List.List{}) {
		return 0, http.StatusBadRequest, fmt.Errorf("List_id = %s doesn't exists", list_id)
	}

	return list_id, http.StatusOK, nil
}

func (h *Handler) listAvailability(context *gin.Context) {
	log.Info().Msg("listAvailability process request.")

	list_id, err := strconv.Atoi(context.Param("list_id"))
	if err != nil {
		log.Error().Msgf("Error during listAvailability. Reason : %s", err.Error())
		initErrorResponce(context, http.StatusBadRequest, err.Error())
		return
	}

	list, err := h.services.GetListById(list_id)
	if err != nil {
		log.Error().Msgf("Error during listAvailability. Reason : %s", err.Error())
		initErrorResponce(context, http.StatusInternalServerError, err.Error())
		return
	}
	if (list == List.List{}) {
		log.Error().Msgf("List doesn't exist.")
		initErrorResponce(context, http.StatusInternalServerError, err.Error())
		return
	} else {
		context.Set("list_id", list_id)
		return
	}

}
