package main

import (
	"fmt"
	"os"
	"runtime"
	"strings"
	"sync"

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
	ch := make(chan error, 3)

	var wg sync.WaitGroup

	ch <- initConfig()

	wg.Add(2)
	go func() {
		defer wg.Done()
		ch <- godotenv.Load()
	}()
	go func() {
		defer wg.Done()
		ch <- initLogger()
	}()

	wg.Wait()
	close(ch)
	for errVal := range ch {
		if errVal != nil {
			panic(fmt.Sprintf("Error during init the application. Reason: %s", errVal.Error()))
		}
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

	service.Audit = service.InitAudit(repository, viper.GetString("audit_level"))

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

func initLogger() error {
	var err error

	if log.Logger, err = server.InitLogger(); err != nil {
		return err
	} else {
		log.Info().Msg("Logger initialized.")
		return nil
	}

}
