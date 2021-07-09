package Router

import (
	"sync"

	"github.com/DimKush/guestbook/tree/main/backend/internal/Contrtollers/Ping"
	"github.com/DimKush/guestbook/tree/main/backend/internal/Logger"
	"github.com/gorilla/mux"
)

type Router interface {
	ReturnRouter() *mux.Router
}

type router struct {
	contrtollers map[string]Ping.Controller
	router       *mux.Router
}

var once sync.Once

func (data *router) init() {
	Logger.Instance().Log().Info().Msgf("Start proccess request init()")
	data.router = mux.NewRouter()
	data.handlersRegist()
}

func (data *router) handlersRegist() {
	// registration
	Logger.Instance().Log().Info().Msgf("Start proccess request Route()")
	data.router.HandleFunc("/main/Ping", Ping.NewPing().Execute).GetError()

}

func (data *router) ReturnRouter() *mux.Router {
	return data.router
}
