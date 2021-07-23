package main

import (
	"fmt"

	server "github.com/DimKush/guestbook/tree/main"
	"github.com/DimKush/guestbook/tree/main/pkg/handler"
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

	if err := server.Run(strPort, routes); err != nil {
		return fmt.Errorf("Cannot run server on port : %s. Reason : %s", strPort, err.Error())
	}

	return nil
}
