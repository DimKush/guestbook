package Logger

import (
	"io/ioutil"
	"log"
	"os"
	"runtime"
	"strings"
	"sync"
	"time"

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
	today_log_file    string
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
	once.Do(func() {
		// check log file
		data.checkLogDateFile()

		if Severity >= data.current_log_level {
			return
		}
		var severityStr string
		switch Severity {
		case ERROR:
			{
				severityStr = "ERROR"
			}
		case WARNING:
			{
				severityStr = "WARNING"
			}
		case INFO:
			{
				severityStr = "INFO"
			}
		case DEBUG:
			{
				severityStr = "DEBUG"
			}
		case TRACE:
			{
				severityStr = "TRACE"
			}
		}

		log.Printf("[%s] %s", severityStr, message)
	})

}

//guestbook_YYYY-MM-DD.log
func (data *logger) createLogNewDate() {
	current_dt := time.Now()

	var strb strings.Builder
	strb.WriteString("guestbook_")

	current_date_str := current_dt.Format("2009-01-31")

	strb.WriteString(current_date_str)

	strb.WriteString(".log")

	data.today_log_file = strb.String()

	strb.Reset()
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
		log.Fatalf("Cannot create file %s", strb.String())
	}

	log_file.Close()
}

func (data *logger) checkLogDateFile() {
	files, err := ioutil.ReadDir("path_to_logs")
	if err != nil {
		log.Fatalf("Cannot open the directory with logs %s", data.path_to_logs)
	}

	todayLogFileExists := false
	for _, file := range files {
		if file.Name() == data.today_log_file {
			todayLogFileExists = true
		}
	}

	if !todayLogFileExists {
		data.createLogNewDate()
	}
}

func (data *logger) Init() {
	strLevel := strings.ToUpper(Configurator.Instance().GetLogLevel())

	switch strLevel {
	case "ERROR":
		{
			data.current_log_level = ERROR
		}
	case "WARNING":
		{
			data.current_log_level = WARNING
		}
	case "INFO":
		{
			data.current_log_level = INFO
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

	// mask of the log file name is YYYY-MM-DD.log
	if Configurator.Instance().GetLogPath() == "" {
		log_path_dir_foo := func() string {
			if runtime.GOOS == "windows" {
				return "c:\\dimkush_guestbook\\log\\"
			}

			return "/opt/dimkush_guestbook/log"
		}
		data.path_to_logs = log_path_dir_foo()
	} else {
		data.path_to_logs = Configurator.Instance().GetLogPath()
	}

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

	file, err := os.OpenFile(strb.String(), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("Cannot open log path %s", strb.String())
	}

	log.SetOutput(file)
}
