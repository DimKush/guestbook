package router

import (
	"sync"

	"github.com/gorilla/mux"
)

type Router interface {
}

type router struct {
	router *mux.Router
}

var once sync.Once
var instance *router = nil

func Instance() Router {
	once.Do(func() {
		if instance == nil {
			instance = new(router)
		}
	})

	return instance
}

func (data *router) init() {

}
