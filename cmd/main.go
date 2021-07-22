package main

import (
	"fmt"

	server "github.com/DimKush/guestbook/tree/main"
)

func main() {
	server := new(server.Server)

	//TODO : from config
	strPort := "8040"
	if err := server.Run(strPort); err != nil {
		panic(fmt.Sprintf("Cannot run server on port : %s. Reason : %s", strPort, err.Error()))
	}
}
