package Alive

import (
	"net/http"

	"github.com/DimKush/guestbook/tree/main/internal/Configurator"
	"github.com/DimKush/guestbook/tree/main/internal/Logger"
	"github.com/DimKush/guestbook/tree/main/internal/utils"
)

type Controller interface {
	Execute(writer http.ResponseWriter, reader http.Request)
}

type Alive struct{}

func (data *Alive) Execute(writer http.ResponseWriter, reader http.Request) {
	Logger.Instance().Log().Debug().Msgf("Execute process request")

	serviceName := "Service " + Configurator.Instance().GetServiceName() + "is alive."

	bytes, err := utils.SendOkResponce(serviceName)

	if err != nil {
		Logger.Instance().Log().Error().Msgf("Error %s", err.Error())
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.Write(bytes)
}

func NewAlive() Controller {
	return &Alive{}
}
