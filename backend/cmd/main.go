package main

import (
	"net/http"
	"time"

	"github.com/DimKush/guestbook/tree/main/backend/internal/Configurator"
	"github.com/DimKush/guestbook/tree/main/backend/internal/Contrtollers/Ping"
	"github.com/DimKush/guestbook/tree/main/backend/internal/Logger"
	"github.com/gorilla/mux"
)

func main() {
	Logger.Instance().Log().Info().Msgf("Starting server %s:%s", Configurator.Instance().GetHost(), Configurator.Instance().GetPort())
	r := mux.NewRouter()

	s := &http.Server{
		Addr:         Configurator.Instance().GetFullAddress(),
		Handler:      r,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
		IdleTimeout:  1 * time.Second,
	}

	var hrd Ping.Ping
	r.HandleFunc("/main/Ping", hrd.Execute)

	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		Logger.Instance().Log().Fatal().Msgf("Error during starting the server %s", err.Error())
	} else {
		Logger.Instance().Log().Info().Msgf("Server starts %v", s)
	}
}

func Hello(w http.ResponseWriter, r *http.Request) {
	Logger.Instance().Log().Info().Msgf("Server Hello %v", r)
}
