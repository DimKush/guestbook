package GetHealth

import (
	"encoding/json"
	"net/http"

	"github.com/DimKush/guestbook/tree/main/backend/internal/Logger"
)

type Ping struct {
	Service_name string
}

func (data *Ping) Execute(writer http.ResponseWriter, reader *http.Request) {
	Logger.Instance().Log().Debug().Msg("Execute process request")

	err := json.NewDecoder(reader.Body).Decode(&data)
	if err != nil {
		Logger.Instance().Log().Error().Msgf("Cannot parse json %s", err.Error())
		return
	}

}
