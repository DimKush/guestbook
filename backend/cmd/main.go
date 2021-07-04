package main

import (
	"github.com/DimKush/guestbook/tree/main/backend/internal/Logger"
)

func main() {
	//Logger.Instance().Write(Logger.ERROR, "Logger init t")
	//Logger.Instance().Write(Logger.DEBUG, "Logger init DEBUG")
	Logger.Instance().Log().Error().Msgf("Logger init t %d", 10)
	Logger.Instance().Log().Debug().Msgf("Logger init t %d", 10)
	Logger.Instance().Log().Info().Msgf("Logger init t %d", 10)

}
