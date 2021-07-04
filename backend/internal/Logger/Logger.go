package Logger

import (
	"fmt"
	"io/ioutil"
	"os"
	"runtime"
	"strings"
	"sync"
	"time"

	"github.com/DimKush/guestbook/tree/main/backend/internal/Configurator"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

const (
	ERROR   = 0
	WARNING = iota
	INFO    = iota
	DEBUG   = iota
	TRACE   = iota
)

type Logger struct {
	current_log_level int
	path_to_logs      string
	today_log_file    string
	inst              *zerolog.Logger
}

var once sync.Once
var onceWrite sync.Once

func (data *Logger) Instance() *zerolog.Logger {
	once.Do(func() {
		if data.inst == nil {
			logger := data.init()
			data.inst = &logger
		}
	})

	return data.inst
}

func (data *Logger) createLogNewDate() {
	var strb strings.Builder

	log_path_dir_foo := func() string {
		if runtime.GOOS == "windows" {
			return "\\"
		}

		return "/"
	}
	var separator string = log_path_dir_foo()

	strb.WriteString(data.path_to_logs)
	strb.WriteString(separator)
	strb.WriteString(data.today_log_file)

	log_file, err := os.Create(strb.String())

	if err != nil {
		log.Fatal().Msg(fmt.Sprintf("Cannot create the file %s", strb.String()))
	}

	log_file.Close()
}

func (data *Logger) checkLogDateFile() bool {
	files, err := ioutil.ReadDir(data.path_to_logs)
	if err != nil {
		log.Fatal().Msg(fmt.Sprintf("Cannot open the directory with logs %s", data.path_to_logs))
	}

	for _, file := range files {
		if file.Name() == data.today_log_file {
			return true
		}
	}

	return false
}

func (data *Logger) init() zerolog.Logger {
	strLevel := strings.ToUpper(Configurator.Instance().GetLogLevel())

	switch strLevel {
	case "ERROR":
		{
			log.Level(zerolog.ErrorLevel)
		}
	case "WARNING":
		{
			log.Level(zerolog.WarnLevel)
		}
	case "INFO":
		{
			log.Level(zerolog.WarnLevel)
		}
	case "DEBUG":
		{
			data.current_log_level = DEBUG
		}
	case "TRACE":
		{
			data.current_log_level = TRACE
		}
	default:
		{
			data.current_log_level = ERROR
		}
	}

	if Configurator.Instance().GetLogPath() == "" {
		log_path_dir_foo := func() string {
			if runtime.GOOS == "windows" {
				return "c:\\dimkush_guestbook\\log\\"
			}

			return "/opt/dimkush_guestbook/log/"
		}
		data.path_to_logs = log_path_dir_foo()
	} else {
		data.path_to_logs = Configurator.Instance().GetLogPath()
	}

	current_dt := time.Now()

	var strb strings.Builder
	strb.WriteString("guestbook_")

	current_date_str := current_dt.Format("2006-Jan-02")

	strb.WriteString(current_date_str)

	strb.WriteString(".log")

	data.today_log_file = strb.String()

	if !data.checkLogDateFile() {
		data.createLogNewDate()
	}

	strb.Reset()

	strb.WriteString(data.path_to_logs)
	strb.WriteString(data.today_log_file)

	file, err := os.OpenFile(strb.String(), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal().Msg(fmt.Sprintf("Cannot open the log file %s", strb.String()))
	}

	log.Output(file)
	//log.SetOutput(file)
	return log.Logger
}
