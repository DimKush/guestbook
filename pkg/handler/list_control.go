package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/DimKush/guestbook/tree/main/internal/entities/List"
	"github.com/gin-gonic/gin"
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
