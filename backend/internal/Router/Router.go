package Router

import (
	"sync"

	"github.com/DimKush/guestbook/tree/main/backend/internal/Contrtollers/Ping"
	"github.com/DimKush/guestbook/tree/main/backend/internal/Logger"
	"github.com/gorilla/mux"
)

type Router interface {
	Route()
	ReturnRouter() *mux.Router
}

type router struct {
	contrtollers map[string]Ping.Controller
	router       *mux.Router
}

var once sync.Once
var instance *router = nil

func Instance() Router {
	once.Do(func() {
		if instance == nil {
			instance = new(router)
			instance.init()
		}
	})

	return instance
}

func (data *router) init() {
	data.router = mux.NewRouter()
}

func (data *router) Route() {
	// registration
	Logger.Instance().Log().Info().Msgf("Start proccess request Route()")

	data.router.HandleFunc("main/Ping", Ping.NewPing().Execute)
}

func (data *router) ReturnRouter() *mux.Router {
	return data.router
}
