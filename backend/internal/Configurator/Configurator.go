package Configurator

import (
	"io/ioutil"
	"log"
	"runtime"
	"sync"

	"gopkg.in/yaml.v2"
)

type Configurator interface {
	//Init() error
	GetLogLevel() string
	GetLogPath() string
}

type configurator struct {
	Host      string `yaml:"host"`
	Port      int    `yaml:"port"`
	Log_level string `yaml:"log_level"`
	Log_path  string `yaml:"log_path"`
}

var instance *configurator = nil
var once sync.Once

func Instance() Configurator {
	once.Do(func() {
		if instance == nil {
			instance = new(configurator)
			instance.init_inside()
		}
	})

	return instance
}

func (data *configurator) init_inside() {
	// a little bit js style

	path_to_conf_foo := func() string {
		if runtime.GOOS == "windows" {
			return "c:\\dimkush_guestbook\\conf\\config.yaml"
		}

		return "/opt/dimkush_guestbook/conf/config.yaml"
	}

	default_path_to_conf := path_to_conf_foo()

	yamlFile, err := ioutil.ReadFile(default_path_to_conf)
	if err != nil {
		log.Fatalf("Cannot find the config file on path %s", default_path_to_conf)
	}

	err = yaml.Unmarshal(yamlFile, data)
	if err != nil {
		log.Fatalf("Cannot unmarshall the config file on path %s", default_path_to_conf)
	}
}

func (data *configurator) GetLogLevel() (level string) {
	return data.Log_level
}

func (data *configurator) GetLogPath() (pathToLog string) {
	return data.Log_path
}
