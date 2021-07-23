package handler

import (
	"github.com/DimKush/guestbook/tree/main/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services service.Service
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api")
	{
		api.GET("/status", h.status)
		lists := router.Group("/lists")
		{
			lists.POST("/", h.createList)
			lists.GET("/", h.getAllLists)
			lists.GET("/:list_id", h.getListById)
			lists.PUT("/:list_id", h.updateListById)
			lists.DELETE("/:list_id", h.dropListById)

			events := router.Group(":id/items")
			{
				events.POST("/", h.createEvent)
				events.GET("/", h.getAllEvents)
				events.GET("/:item_id", h.getEventById)
				events.PUT("/:item_id", h.updateEventById)
				events.DELETE("/:item_id", h.dropEventById)
			}
		}
	}

	return router
}
