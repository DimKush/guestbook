package Configurator

import (
	"fmt"
	"io/ioutil"
	"log"
	"runtime"
	"strings"
	"sync"

	"gopkg.in/yaml.v2"
)

type Configurator interface {
	//Init() error
	GetFullAddress() string
	GetLogLevel() string
	GetLogPath() string
	GetPort() string
}

type configurator struct {
	Port      string `yaml:"port"`
	Log_level string `yaml:"log_level"`
	Log_path  string `yaml:"log_path"`
	Database  struct {
		Db_name     string `yaml:"db_name"`
		Db_core     string `yaml:"db_core"`
		Db_user     string `yaml:"db_user"`
		Db_password string `yaml:"db_password"`
		Db_port     string `yaml:"db_port"`
	}
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

	fmt.Printf("%v", data)
}

func (data *configurator) GetLogLevel() (level string) {
	return data.Log_level
}

func (data *configurator) GetLogPath() (pathToLog string) {
	return data.Log_path
}

func (data *configurator) GetPort() string {
	return data.Port
}

func (data *configurator) GetFullAddress() string {
	var strb strings.Builder
	strb.WriteString(":")
	strb.WriteString(data.Port)

	return strb.String()
}

/*

func (data *configurator) GetDbName() string {
	return data.Db_conf.Db_name
}

func (data *configurator) GetDbCore() string {
	return data.Db_conf.Db_core
}

func (data *configurator) GetDbUser() string {
	return data.Db_conf.Db_user
}

func (data *configurator) GetDbPassword() string {
	return data.Db_conf.Db_password
}

func (data *configurator) GetDbPort() string {
	return data.Db_conf.Db_port
}
*/
