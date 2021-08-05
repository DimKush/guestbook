package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type ErrorResponce struct {
	Status  string `json:"Status"`
	Message string `json:"Message"`
}

func initErrorResponce(c *gin.Context, statusCode int, errMessage string) {
	log.Error().Msgf("Json Error responce with message %s", errMessage)
	c.AbortWithStatusJSON(statusCode, ErrorResponce{Status: "Error", Message: errMessage})
}

func initOkResponce(c *gin.Context, params map[string]interface{}) {
	params["Status"] = "OK"

	c.JSON(http.StatusOK, params)
}
