package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/DimKush/guestbook/tree/main/internal/entities/List"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func (h *Handler) createList(context *gin.Context) {
	log.Info().Msg("Handler createList process request.")

	var newList List.List

	if err := context.BindJSON(&newList); err != nil {
		initErrorResponce(context, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.CreateList(newList); err != nil {
		initErrorResponce(context, http.StatusInternalServerError, err.Error())
		return
	}

	initOkResponce(context, map[string]interface{}{})
}

func (h *Handler) getAllLists(context *gin.Context) {
	lists, err := h.services.GetAllLists()
	if err != nil {
		initErrorResponce(context, http.StatusInternalServerError, err.Error())
		return
	}

	initOkResponce(context, map[string]interface{}{
		"Result": lists,
	})
}

func (h *Handler) getListsByParams(context *gin.Context) {
	log.Info().Msg("Handler GetListsByParams process request.")
	var listsParams List.List

	if err := context.BindJSON(&listsParams); err != nil {
		//listsParams = List.List{}
		initErrorResponce(context, http.StatusBadRequest, err.Error())
		return
	}

	lists, err := h.services.GetListsByParams(listsParams)

	if err != nil {
		initErrorResponce(context, http.StatusInternalServerError, err.Error())
		return
	}

	if len(lists) == 0 {
		err := fmt.Errorf("Didn't find anything.")
		initErrorResponce(context, http.StatusOK, err.Error())
	}

	initOkResponce(context, map[string]interface{}{
		"Result": lists,
	})
}

func (h *Handler) getListById(context *gin.Context) {
	log.Info().Msg("Handler getListById process request.")

	list_id, err := strconv.Atoi(context.Param("list_id"))
	if err != nil {
		initErrorResponce(context, http.StatusBadRequest, err.Error())
		return
	}

	list, err := h.services.GetListById(list_id)
	if err != nil {
		initErrorResponce(context, http.StatusInternalServerError, err.Error())
		return
	}

	initOkResponce(context, map[string]interface{}{
		"Result": list,
	})
}

func (h *Handler) updateListById(context *gin.Context) {
	log.Info().Msg("Handler updateEventById process request.")

	id, http_code, err := h.ControlListExist(context)
	if err != nil {
		initErrorResponce(context, http_code, err.Error())
		return
	}

	// unmarshall json
	var lst List.List

	if err := context.BindJSON(&lst); err != nil {
		initErrorResponce(context, http.StatusInternalServerError, "Internal server error.")
		return
	}

	lst.Id = id

	// if h.services.Up

	// context.Status(http.StatusOK)
}

func (h *Handler) dropListById(context *gin.Context) {
	log.Info().Msg("Handler dropListById process request.")

	list_id, err := strconv.Atoi(context.Param("list_id"))
	if err != nil {
		initErrorResponce(context, http.StatusBadRequest, err.Error())
		return
	}

	if err = h.services.DeleteListById(list_id); err != nil {
		initErrorResponce(context, http.StatusInternalServerError, err.Error())
		return
	}

	initOkResponce(context, map[string]interface{}{
		"Message": fmt.Sprintf("List (id = %d) has been deleted.", list_id),
	})
}

func (h *Handler) getAutoListId(context *gin.Context) {
	id, err := h.services.GetAutoListId()

	if err != nil {
		initErrorResponce(context, http.StatusInternalServerError, err.Error())
		return
	}

	initOkResponce(context, map[string]interface{}{
		"Result": id,
	})
}
