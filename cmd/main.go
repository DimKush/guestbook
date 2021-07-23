package main

import (
	"fmt"
	"runtime"
	"strings"

	server "github.com/DimKush/guestbook/tree/main"
	"github.com/DimKush/guestbook/tree/main/pkg/handler"
	"github.com/spf13/viper"
)

func main() {
	server := new(server.Server)
	if err := run(server); err != nil {
		panic(err.Error())
	}

}

func run(server *server.Server) error {
	//TODO : from config
	strPort := "8040"
	handler := new(handler.Handler)

	routes := handler.InitRoutes()

	// read config
	if err := InitConfig(); err != nil {
		return fmt.Errorf("Cannot read service config. Reason: %s", err.Error())
	}

	if err := server.Run(strPort, routes); err != nil {
		return fmt.Errorf("Cannot run server on port : %s. Reason : %s", strPort, err.Error())
	}

	return nil
}

func InitConfig() error {
	platform := strings.ToLower(runtime.GOOS)

	var confDirPath string
	switch platform {
	case "linux":
		{
			confDirPath = "/opt/dimkush_guestbook/conf"
		}
	case "windows":
		{
			confDirPath = "c:\\dimkush_guestbook\\conf"
		}
	}

	viper.AddConfigPath(confDirPath)
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
