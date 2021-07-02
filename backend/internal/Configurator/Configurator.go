package Configurator

import (
	"os"
)

type configurator struct {
	host     int
	port     int
	logLevel string
}

func (data *configurator) Init() (status error) {
	default_path_to_conf := "/opt/dimkush_guestbook/conf/config.yaml"

	_, err := os.Stat(default_path_to_conf)
	if err != nil {
		Logger.Write(Logger.ERROR, err.Error())
	}
}

var Configurator = configurator{}
C