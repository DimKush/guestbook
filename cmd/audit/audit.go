package main

import (
	"log"
	"net/http"
	"time"

	"github.com/DimKush/guestbook/tree/main/internal/Configurator"
	"github.com/DimKush/guestbook/tree/main/internal/Contrtollers/main/Ping"
	"github.com/DimKush/guestbook/tree/main/internal/Logger"
	"github.com/gorilla/mux"
)

func registerRouter() *mux.Router {
	router := new(mux.Router)

	router.HandleFunc("/audit/Ping", Ping.NewPing().Execute).GetError()

	return router
}

func main() {
	Configurator.Instance().Init("audit")

	port, err := Configurator.Instance().GetPort()
	if err != nil {
		log.Fatalf(err.Error())
	}

	addrStr := "localhost" + ":" + port
	Logger.Instance().Log().Info().Msgf("Starting server localhost:%s", addrStr)

	s := &http.Server{
		Addr:         addrStr,
		Handler:      registerRouter(),
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
		IdleTimeout:  1 * time.Second,
	}

	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		Logger.Instance().Log().Fatal().Msgf("Error during starting the server %s", err.Error())
	} else {
		Logger.Instance().Log().Info().Msgf("Server starts %v", s)
	}

}
