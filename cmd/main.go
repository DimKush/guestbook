package main

import (
	"fmt"

	server "github.com/DimKush/guestbook/tree/main"
)

func main() {
	server := new(server.Server)
	if err := run(server); err != nil {
		panic(err.Error())
	}
	//TODO : from config

}

func run(server *server.Server) error {
	strPort := "8040"
	if err := server.Run(strPort); err != nil {
		panic(fmt.Sprintf("Cannot run server on port : %s. Reason : %s", strPort, err.Error()))
	}
}
