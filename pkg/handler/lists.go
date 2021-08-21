package handler

import (
	"fmt"
	"net/http"

	"github.com/DimKush/guestbook/tree/main/internal/entities/List"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func (h *Handler) createList(context *gin.Context) {
	id, _ := context.Get(userCTX)
	context.JSON(http.StatusOK, map[string]interface{}{
		userCTX: id,
	})

}

func (h *Handler) getAllLists(context *gin.Context) {
	lists, err := h.services.GetAllLists()
	if err != nil {
		initErrorResponce(context, http.StatusInternalServerError, err.Error())
	}

	initOkResponce(context, map[string]interface{}{
		"Result": lists,
	})
}

func (h *Handler) GetListsByParams(context *gin.Context) {
	log.Info().Msg("Handler GetListsByParams process request.")
	var listsParams List.List

	if err := context.BindJSON(&listsParams); err != nil {
		listsParams = List.List{}
	}

	lists, err := h.services.GetListsByParams(listsParams)

	if err != nil {
		initErrorResponce(context, http.StatusInternalServerError, err.Error())
	}

	if len(lists) == 0 {
		err := fmt.Errorf("Didn't find anything.")
		initErrorResponce(context, http.StatusOK, err.Error())

		return
	}

	initOkResponce(context, map[string]interface{}{
		"Result": lists,
	})
}

func (h *Handler) GetAllUsernames(context *gin.Context) {
	log.Info().Msg("Handler GetAllUsernames process request.")

	users, err := h.services.GetAllUsernames()

	if err != nil {
		initErrorResponce(context, http.StatusInternalServerError, err.Error())
	}

	initOkResponce(context, map[string]interface{}{
		"Result": users,
	})
}

func (h *Handler) getListById(context *gin.Context) {

}

func (h *Handler) updateListById(context *gin.Context) {

}

func (h *Handler) dropListById(context *gin.Context) {

}
