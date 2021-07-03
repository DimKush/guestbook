package main

import (
	"github.com/DimKush/guestbook/tree/main/backend/internal/Logger"
)

func main() {
	Logger.Instance().Write(Logger.ERROR, "Logger init t")
	Logger.Instance().Write(Logger.DEBUG, "Logger init DEBUG")
}
