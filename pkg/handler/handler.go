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
	config.AllowMethods = []string{"POST", "GET", "DELETE", "PUT"}
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
			lists.DELETE("/:list_id", h.dropListById) // TODO : Need to delete items too

			events := lists.Group(":list_id/items", h.getListById)
			{
				events.GET("/availability", h.getItemsAvailability)
				events.POST("/params", h.GetItemsByParams)
				events.POST("/types", h.GetItemsTypes)
				events.POST("/create", h.createEvent)
				events.GET("/", h.getAllItemsByListId)
				events.GET("/:item_id", h.getItemById)
				events.PUT("/:item_id", h.updateItemById)
				events.DELETE("/:item_id", h.deleteItemById)
			}

			lists.POST("/items/all", h.getAllUsersEvents)
		}
	}

	return router
}

func HandlerInit(service *service.Service) *Handler {
	return &Handler{
		services: *service,
	}
}
