package Ping

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
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
	Service_port int    `json:"service_port"`
}

type AliveAnswer struct {
	Status string `json:"Status"`
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

	if data.Service_name == "" || data.Service_port == 0 {
		bytes, err := utils.SenErrorMessage("Error, Incorrect input params", err.Error())

		if err != nil {
			Logger.Instance().Log().Error().Msg(err.Error())
		}

		writer.Write(bytes)
		return
	}
	//TODO build request string
	getStr := "localhost:" + strconv.Itoa(data.Service_port) + "/" + data.Service_name + "/Alive"

	fmt.Println(getStr)
	resp, err := http.Get(getStr)
	respData := AliveAnswer{}

	fmt.Println("1")
	err = json.NewDecoder(resp.Body).Decode(&respData)

	fmt.Println("2")
	if respData.Status == "OK" {
		bytes, _ := utils.SendOkResponce(fmt.Sprintf("Service %s is alive", data.Service_name))

		writer.Write(bytes)
	}
	fmt.Println("3")
}

func NewPing() Controller {
	return &Ping{}
}
