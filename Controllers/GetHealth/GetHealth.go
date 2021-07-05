package GetHealth

import (
	"encoding/json"
	"net/http"

	"github.com/DimKush/guestbook/tree/main/backend/internal/Logger"
)

type GetHealth struct {
	service_name string
}

func (data *GetHealth) Execute(writer http.ResponseWriter, reader *http.Request) {
	Logger.Instance().Log().Debug().Msg("Execute process request")

	err := json.NewDecoder(reader.Body).Decode(&data)
	if err != nil {
		Logger.Instance().Log().Error().Msgf("Cannot parse json %s", err.Error())
		return
	}

	Logger.Instance().Log().Debug().Msgf("%v", data)
}
