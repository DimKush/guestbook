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

const (
	DB_POSTGRES = iota
	DB_MYSQL    = iota
	DB_SQLITE   = iota
)

type Configurator interface {
	//Init() error
	GetFullAddress() string
	GetLogLevel() string
	GetLogPath() string
	GetPort(string) string
	GetDbConnectGorm(core int) string
	GetDbConnectionPool() int
}

type configurator struct {
	main struct {
		Port      string `yaml:"port"`
		Log_level string `yaml:"log_level"`
		Log_path  string `yaml:"log_path"`
	}
	audit struct {
		Port        string `yaml:"port"`
		Audit_level string `yaml:"audit_level"`
		Log_level   string `yaml:"log_level"`
	}
	database struct {
		Db_name             string `yaml:"db_name"`
		Db_core             string `yaml:"db_core"`
		Db_user             string `yaml:"db_user"`
		Db_password         string `yaml:"db_password"`
		Db_port             string `yaml:"db_port"`
		Db_connections_pool int    `yaml:"db_connections_pool"`
	}

	serviceName string
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

func (data *configurator) Init(service string) {
	data.serviceName = strings.ToLower(service)

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

	switch data.serviceName {

	case "main":
		{
			return data.main.Log_level
		}
	case "audit":
		{
			return data.audit.Log_level
		}

	default:
		{
			return string("")
		}
	}
}

func (data *configurator) GetLogPath() (pathToLog string) {
	return data.main.Log_path
}

func (data *configurator) GetPort(service string) string {
	serviceStr := strings.ToUpper(service)

	switch serviceStr {
	case "main":
		{

		}
	}

}

func (data *configurator) GetFullAddress() string {
	var strb strings.Builder
	strb.WriteString(":")
	strb.WriteString(data.Port)

	return strb.String()
}

func (data *configurator) GetDbConnectGorm(core int) string {
	var gormStr string
	switch core {
	case DB_POSTGRES:
		{
			gormStr = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
				"localhost",
				data.database.Db_port,
				data.database.Db_user,
				data.database.Db_name,
				data.database.Db_password,
			)
		}
	default:
		{
			//panic()
		}
	}

	return gormStr
}

func (data *configurator) GetDbConnectionPool() int {
	return data.database.Db_connections_pool
}
