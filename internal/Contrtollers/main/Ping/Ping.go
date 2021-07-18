package Ping

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/DimKush/guestbook/tree/main/internal/AuditProxy"
	"github.com/DimKush/guestbook/tree/main/internal/Logger"
	"github.com/DimKush/guestbook/tree/main/internal/utils"
)

type Controller interface {
	Execute(writer http.ResponseWriter, reader *http.Request)
}

type Ping struct {
	Service_name string `json:"service_name"`
	Service_port string `json:"service_port"`
}

func (data *Ping) Execute(writer http.ResponseWriter, reader *http.Request) {
	Logger.Instance().Log().Info().Msg("Execute process request")

	// TODO : for test
	// TODO : need a parallel execution
	AuditProxy.WriteEvent("debug", time.Now(), "main", false, "test descr")

	writer.Header().Set("Content-Type", "application/json")
	err := json.NewDecoder(reader.Body).Decode(&data)

	if err != nil {
		strErr := fmt.Sprintf("Cannot parse json %s", err.Error())
		Logger.Instance().Log().Error().Msg(strErr)
		bytes, err := utils.SenErrorMessage("Error during execute Ping function", err.Error())

		if err != nil {
			Logger.Instance().Log().Error().Msg(err.Error())
		}

		writer.Write(bytes)
		return
	}

	if data.Service_name == "" || data.Service_port == "" {
		bytes, err := utils.SenErrorMessage("Error, Incorrect input params", err.Error())

		if err != nil {
			Logger.Instance().Log().Error().Msg(err.Error())
		}

		writer.Write(bytes)
		return
	}
	// build request string
	//getStr := "localhost:"
}

func NewPing() Controller {
	return &Ping{}
}
