package logging

import (
	"log"
	"os"
)

func NewDefaultLogger(level LogLevel) *DefaultLogger {
	flags := log.Ltime | log.Lmsgprefix | log.Ldate | log.LUTC
	return &DefaultLogger{
		minLevel: level,
		loggers: map[LogLevel]*log.Logger{
			Trace:       log.New(os.Stdout, "TRACE ", flags),
			Debug:       log.New(os.Stdout, "DEBUG ", flags),
			Information: log.New(os.Stdout, "INFO ", flags),
			Warning:     log.New(os.Stdout, "WARN ", flags),
			Fatal:       log.New(os.Stdout, "FATAL ", flags),
		},
		triggerPanic: true,
	}
}
