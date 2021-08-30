package handler

import (
	"github.com/DimKush/guestbook/tree/main/pkg/service"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services service.Service
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	// CORS configuration
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	config.AllowHeaders = []string{"Content-type", "Authorization"}
	config.AllowCredentials = true

	router.Use(cors.New(config))

	status := router.GET("", h.status)
	status.GET("/status", h.status)

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
		auth.GET("/user", h.userIdentityUsername)
		auth.GET("/logout", h.logout)
	}

	api := router.Group("/api", h.userIdentity)
	{
		users := api.Group("/users")
		{
			users.GET("GetAllUsernames", h.getAllUsernames)
		}

		lists := api.Group("/lists")
		{
			lists.POST("/create", h.createList)
			lists.GET("/", h.getAllLists)
			lists.POST("/params", h.getListsByParams)
			lists.GET("/:list_id", h.getListById)
			lists.GET("/GetAutoListId", h.getAutoListId)
			lists.PUT("/:list_id", h.updateListById)
			lists.DELETE("/:list_id", h.dropListById)

			lists.POST("/items/params", h.getEventsByParams)
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

func HandlerInit(service *service.Service) *Handler {
	return &Handler{
		services: *service,
	}
}
