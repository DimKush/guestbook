package Logger

import (
	"strings"
	"sync"

	"github.com/DimKush/guestbook/tree/main/backend/internal/Configurator"
)

const (
	ERROR   = 0
	WARNING = iota
	INFO    = iota
	DEBUG   = iota
	TRACE   = iota
)

type Logger interface {
	Init()
	Write(Severity int, message string)
}

type logger struct {
	current_log_level int
	path_to_logs      string
}

var instance *logger = nil
var once sync.Once

func Instance() Logger {
	once.Do(func() {
		if instance == nil {
			instance = new(logger)
			instance.Init()
		}
	})

	return instance
}

func (data *logger) Write(Severity int, message string) {
	if Severity > instance.current_log_level {
		return
	} else {

	}
}

func (data *logger) Init() {
	strLevel := strings.ToUpper(Configurator.Instance().GetLogLevel())

	switch strLevel {
	case "ERROR":
		{
			instance.current_log_level = ERROR
		}
	case "WARNING":
		{
			instance.current_log_level = WARNING
		}
	case "INFO":
		{
			instance.current_log_level = INFO
		}
	case "DEBUG":
		{
			instance.current_log_level = DEBUG
		}
	case "TRACE":
		{
			instance.current_log_level = TRACE
		}
	default:
		{
			instance.current_log_level = ERROR
		}
	}

	// mask of the log file name is YYYY-MM-DD.log
	if Configurator.Instance().GetLogPath() == "" {
		instance.path_to_logs = "/opt/dimkush_guestbook/log/"
	} else {
		instance.path_to_logs = Configurator.Instance().GetLogPath()
	}
}
