package main

import (
	"fmt"
	"runtime"
	"strings"

	server "github.com/DimKush/guestbook/tree/main"
	"github.com/DimKush/guestbook/tree/main/pkg/handler"
	"github.com/DimKush/guestbook/tree/main/pkg/repository"
	"github.com/DimKush/guestbook/tree/main/pkg/service"
	"github.com/spf13/viper"
)

func init() {
	// read config
	if err := InitConfig(); err != nil {
		panic(fmt.Sprintf("Cannot read service config. Reason: %s", err.Error()))
	}
}

func main() {
	server := new(server.Server)
	if err := run(server); err != nil {
		panic(err.Error())
	}

}

func run(server *server.Server) error {
	//TODO : from config
	strPort := "8040"

	db_config := repository.Config{
		Host:     viper.GetString("database.host"),
		Port:     viper.GetString("database.port"),
		Username: viper.GetString("database.username"),
		Password: viper.GetString("database.password"),
		Dbname:   viper.GetString("database.dbname"),
		Timezone: viper.GetString("database.timezone"),
		SSLMode:  viper.GetString("database.sslmode"),
	}

	db, err := repository.NewPostgresConnection(db_config)
	if err != nil {
		return fmt.Errorf("Cannot create db connection %v.\nReason: %s", db_config, err.Error())
	}

	repository := repository.RepositoryInit(db)
	services := service.ServiceInit(repository)
	handlers := handler.HandlerInit(services)

	if err := server.Run(strPort, handlers.InitRoutes()); err != nil {
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
