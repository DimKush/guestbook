package Logger

/* TODO comment until new configurator will be done
const (
	ERROR = iota
	WARNING
	INFO
	DEBUG
	TRACE
)

type Logger interface {
	Log() *zerolog.Logger
}

type log_struct struct {
	current_log_level int
	path_to_logs      string
	today_log_file    string
	zLog              *zerolog.Logger
}

var once sync.Once
var onceWrite sync.Once
var instance *log_struct = nil
var logger_data *log_struct = nil

func Instance() Logger {
	once.Do(func() {
		if instance == nil {
			instance = new(log_struct)
			instance.init()
		}
	})

	return instance
}

func (data *log_struct) Log() *zerolog.Logger {
	return data.zLog
}

func (data *log_struct) createLogNewDate() {
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
		log.Fatalf("Cannot create the file %s", strb.String())
	}

	log_file.Close()
}

func (data *log_struct) checkLogDateFile() bool {
	files, err := ioutil.ReadDir(data.path_to_logs)
	if err != nil {
		log.Fatalf("Cannot open the directory with logs %s", data.path_to_logs)
	}

	for _, file := range files {
		if file.Name() == data.today_log_file {
			return true
		}
	}

	return false
}

func (data *log_struct) init() {
	strLevel := strings.ToUpper(Configurator.Instance().GetLogLevel())

	var complete_path = func(data string) string {
		var strb strings.Builder
		if runtime.GOOS == "windows" {
			strb.WriteString(data)
			strb.WriteString("\\")
			return strb.String()
		} else {
			strb.WriteString(data)
			strb.WriteString("/")
			return strb.String()
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
		data.path_to_logs = complete_path(Configurator.Instance().GetLogPath())
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
		log.Fatalf("Cannot open the log file %s", strb.String())
	}

	zlogger := zerolog.New(file).With().Caller().Timestamp().Logger().Output(file)
	data.zLog = &zlogger

	switch strLevel {
	case "ERROR":
		{
			zerolog.SetGlobalLevel(zerolog.ErrorLevel)
		}
	case "WARNING":
		{
			zerolog.SetGlobalLevel(zerolog.WarnLevel)
		}
	case "INFO":
		{
			zerolog.SetGlobalLevel(zerolog.InfoLevel)
		}
	case "DEBUG":
		{
			zerolog.SetGlobalLevel(zerolog.DebugLevel)
		}
	case "TRACE":
		{
			zerolog.SetGlobalLevel(zerolog.TraceLevel)
		}
	default:
		{
			zerolog.SetGlobalLevel(zerolog.ErrorLevel)
		}
	}
}
*/
