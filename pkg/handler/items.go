package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/DimKush/guestbook/tree/main/internal/entities/Item"
	"github.com/DimKush/guestbook/tree/main/internal/entities/UserIn"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func (h *Handler) GetItemsByParams(context *gin.Context) {
	log.Info().Msg("Handler GetItemsByParams process request.")

	list_id, err := strconv.Atoi(context.Param("list_id"))
	if err != nil {
		initErrorResponce(context, http.StatusBadRequest, "Incorrect list_id in the url.")
		return
	}
	if list_id == 0 {
		initErrorResponce(context, http.StatusBadRequest, "In the url list_id cannot be 0.")
		return
	}

	// get current user
	h.userIdentity(context)
	var user UserIn.UserIn
	if userId, exists := context.Get(userCTX); exists {
		if convertedId, ok := userId.(int); !ok {
			err := fmt.Errorf("Error parsing userId %v", userId)
			log.Error().Msgf("Error during parsing userIdentity : %s", err.Error())
			initErrorResponce(context, http.StatusInternalServerError, err.Error())
		} else {
			user.Id = convertedId
		}
	} else {
		err := fmt.Errorf("Incorrect current username.")
		log.Error().Msgf("Error during parsing json : %s", err.Error())
		initErrorResponce(context, http.StatusInternalServerError, err.Error())
		return
	}

	// parse json
	var item Item.Item
	if err := context.BindJSON(&item); err != nil {
		log.Error().Msgf("Error during parsing json : %s", err.Error())
		initErrorResponce(context, http.StatusBadRequest, err.Error())
		return
	}

	item.ListId = list_id
	item.ItemOwnerId = user.Id

	items, err := h.services.GetItemsByParams(item)
	if err != nil {
		initErrorResponce(context, http.StatusInternalServerError, "")
		return
	}

	initOkResponce(context, map[string]interface{}{
		"Result": items,
	})
}

func (h *Handler) getAllUsersEvents(context *gin.Context) {
	log.Info().Msg("Handler getAllUsersEvents process request.")

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
	var item Item.Item
	if err := context.BindJSON(&item); err != nil {
		log.Error().Msgf("Error during parsing json : %s", err.Error())
		initErrorResponce(context, http.StatusInternalServerError, err.Error())
		return
	}

	item.ItemOwnerId = user.Id

	items, err := h.services.GetItemsByParams(item)
	if err != nil {
		initErrorResponce(context, http.StatusInternalServerError, "")
		return
	}

	initOkResponce(context, map[string]interface{}{
		"Result": items,
	})
}

func (h *Handler) createEvent(context *gin.Context) {
	log.Info().Msg("Handler createEvent process request.")

	var list_id int
	if id, status, err := h.ControlListExist(context); err != nil {
		initErrorResponce(context, status, err.Error())
		return
	} else {
		list_id = id
	}

	var item Item.Item
	if err := context.BindJSON(&item); err != nil {
		initErrorResponce(context, http.StatusInternalServerError, "Server error.")
		return
	}

	item.ListId = list_id

	if err := h.services.CreateNewItem(item); err != nil {
		initErrorResponce(context, http.StatusInternalServerError, err.Error())
		return
	}

	initOkResponce(context, map[string]interface{}{})
}

func (h *Handler) GetItemsTypes(context *gin.Context) {
	if _, status, err := h.ControlListExist(context); err != nil {
		initErrorResponce(context, status, err.Error())
	}

	var item_type Item.ItemType
	if err := context.BindJSON(&item_type); err != nil {
		initErrorResponce(context, http.StatusInternalServerError, "Server error.")
		return
	}

	types, err := h.services.GetItemTypesByParams(item_type)
	if err != nil {
		initErrorResponce(context, http.StatusInternalServerError, "Server error.")
		return
	}

	initOkResponce(context, map[string]interface{}{
		"Result": types,
	})
}

func (h *Handler) getItemsAvailability(context *gin.Context) {
	log.Info().Msg("Handler getItemsCount process request.")

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

	list_id, err := strconv.Atoi(context.Param("list_id"))
	if err != nil {
		initErrorResponce(context, http.StatusBadRequest, err.Error())
		return
	}

	items_count, err := h.services.GetItemsAvailability(list_id, user)
	if err != nil {
		initErrorResponce(context, http.StatusInternalServerError, err.Error())
		return
	}

	initOkResponce(context, map[string]interface{}{
		"Count": items_count,
	})
}

func (h *Handler) getAllItemsByListId(context *gin.Context) {
	log.Info().Msg("Handler getAllItems process request.")

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

	list_id, err := strconv.Atoi(context.Param("list_id"))

	if err != nil {
		initErrorResponce(context, http.StatusBadRequest, err.Error())
		return
	}

	var item Item.Item = Item.Item{ListId: list_id, ItemOwnerId: user.Id}

	items, err := h.services.GetItemsByParams(item)
	if err != nil {
		initErrorResponce(context, http.StatusBadRequest, err.Error())
		return
	}

	initOkResponce(context, map[string]interface{}{
		"Result": items,
	})
}

func (h *Handler) getItemById(context *gin.Context) {
	log.Info().Msg("Handler getItemById process request.")

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

	// get item id
	item_id, err := strconv.Atoi(context.Param("item_id"))

	if err != nil {
		initErrorResponce(context, http.StatusBadRequest, err.Error())
		return
	}

	if itemRes, err := h.services.GetItemById(item_id); err != nil {
		initErrorResponce(context, http.StatusBadRequest, err.Error())
		return
	} else {
		initOkResponce(context, map[string]interface{}{
			"Result": itemRes,
		})
		return
	}

}

func (h *Handler) updateEventById(context *gin.Context) {

}

func (h *Handler) dropEventById(context *gin.Context) {

}
