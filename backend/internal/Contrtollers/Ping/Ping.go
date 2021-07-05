package Ping

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/DimKush/guestbook/tree/main/backend/internal/Logger"
	"github.com/DimKush/guestbook/tree/main/backend/internal/utils"
)

type Ping struct {
	Service_name string
}

func (data *Ping) Execute(writer http.ResponseWriter, reader *http.Request) {
	Logger.Instance().Log().Info().Msg("Execute process request")

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

	} else {
		bytes, err := utils.SendOkResponce("Service is online")
		if err != nil {
			Logger.Instance().Log().Error().Msg(err.Error())
		}
		writer.Write(bytes)
	}

}
