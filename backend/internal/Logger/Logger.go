package Logger

const (
	ERROR   = 0
	WARNING = iota
	INFO    = iota
	DEBUG   = iota
	TRACE   = iota
)

type logger struct {
	level string
}

func (data *logger) init() {
	//Co
}

func Write(Severity int, message string) {
	if logger_obj == (logger{}) {
		logger_obj.init()
	}
}

var logger_obj = logger{}
