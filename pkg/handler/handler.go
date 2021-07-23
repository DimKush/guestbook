package handler

import (
	"github.com/DimKush/guestbook/tree/main/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services service.Service
	auth     authHandler
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signIn())
		auth.POST("/sign-in")
	}

	api := router.Group("/api")
	{
		api.GET("/status")
		lists := router.Group("/lists")
		{
			lists.POST("/")
			lists.GET("/")
			lists.GET("/:id")
			lists.PUT("/:id")
			lists.DELETE("/:id")

			events := router.Group(":id/items")
			{
				events.POST("/")
				events.GET("/")
				events.GET("/:item_id")
				events.PUT("/:item_id")
				events.DELETE("/:item_id")
			}
		}
	}

	return router
}
