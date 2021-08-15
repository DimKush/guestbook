package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) createList(context *gin.Context) {
	id, _ := context.Get(userCTX)
	context.JSON(http.StatusOK, map[string]interface{}{
		userCTX: id,
	})

}

func (h *Handler) getAllLists(context *gin.Context) {
	fmt.Println("HERE")
	lists, err := h.services.GetAllLists()
	if err != nil {
		initErrorResponce(context, http.StatusInternalServerError, err.Error())
	}

	initOkResponce(context, map[string]interface{}{
		"Result": lists,
	})
}

func (h *Handler) getListById(context *gin.Context) {

}

func (h *Handler) updateListById(context *gin.Context) {

}

func (h *Handler) dropListById(context *gin.Context) {

}
