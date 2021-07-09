package audit

import (
	"net/http"
	"time"

	"github.com/DimKush/guestbook/tree/main/backend/internal/Configurator"
	"github.com/DimKush/guestbook/tree/main/backend/internal/Logger"
	"github.com/DimKush/guestbook/tree/main/backend/internal/Router"
)

func main() {
	Logger.Instance().Log().Info().Msgf("Starting server localhost:%s", Configurator.Instance().GetPort("audit"))

	s := &http.Server{
		Addr:         Configurator.Instance().GetFullAddress(),
		Handler:      Router.Instance().ReturnRouter(),
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
