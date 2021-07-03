package Configurator

import (
	"log"
	"os"
	"runtime"
	"sync"
)

type Configurator interface {
	Init() error
	GetLogLevel() string
	GetLogPath() string
}

type configurator struct {
	host      int
	port      int
	log_level string
	log_path  string
}

var instance *configurator = nil
var once sync.Once

func Instance() Configurator {
	once.Do(func() {
		if instance == nil {
			instance = new(configurator)
		}
	})

	return instance
}

func (data *configurator) Init() (status error) {
	// a little bit js style

	path_to_conf_foo := func() string {
		if runtime.GOOS == "windows" {
			return "c:\\dimkush_guestbook\\conf\\config.yaml"
		}

		return "/opt/dimkush_guestbook/conf/config.yaml"
	}

	default_path_to_conf := path_to_conf_foo()

	_, err := os.Stat(default_path_to_conf)
	if err != nil {
		log.Fatalf("Cannot find the config file on path %s", default_path_to_conf)
	}

	return nil
}

func (data *configurator) GetLogLevel() (level string) {
	return data.log_level
}

func (data *configurator) GetLogPath() (pathToLog string) {
	return data.log_path
}
