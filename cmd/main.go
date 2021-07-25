package main

import (
	"fmt"
	"os"
	"runtime"
	"strings"

	server "github.com/DimKush/guestbook/tree/main"
	"github.com/DimKush/guestbook/tree/main/pkg/handler"
	"github.com/DimKush/guestbook/tree/main/pkg/repository"
	"github.com/DimKush/guestbook/tree/main/pkg/service"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

func init() {
	// read config
	if err := initConfig(); err != nil {
		panic(fmt.Sprintf("Cannot read service config. Reason: %s", err.Error()))
	}

	// read environment variables
	if err := godotenv.Load(); err != nil {
		panic(fmt.Sprintf("Cannot load environment variables. Reason:%s", err.Error()))
	}

	// logger settings
	if err := loggerInit(); err != nil {
		panic(fmt.Sprintf("Cannot init logger. Reason:%s", err.Error()))
	}

}

func main() {
	server := new(server.Server)
	if err := run(server); err != nil {
		panic(err.Error())
	}

}

func run(server *server.Server) error {
	db_config := repository.Config{
		Host:     viper.GetString("database.host"),
		Port:     viper.GetString("database.port"),
		Username: viper.GetString("database.username"),
		Password: os.Getenv("PG_DB_PASSWORD"),
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

	if err := server.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		return fmt.Errorf("Cannot run server on port : %s. Reason : %s", viper.GetString("port"), err.Error())
	}

	return nil
}

func initConfig() error {
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

func loggerInit() error {
	log.Logger = server.InitLogger()
	log.Error().Msg("Error")

	return nil
}
