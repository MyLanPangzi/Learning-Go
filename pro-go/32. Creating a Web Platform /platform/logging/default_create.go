package logging

import (
	"log"
	"os"
	"platform/config"
	"strings"
)

func NewDefaultLogger(cfg config.Configuration) *DefaultLogger {
	flags := log.Ltime | log.Lmsgprefix | log.Ldate | log.LUTC
	return &DefaultLogger{
		minLevel: getLevel(cfg.GetStringDefault("logging:level", "info")),
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

func getLevel(value string) LogLevel {
	switch strings.ToLower(value) {
	case "trace":
		return Trace
	case "debug":
		return Debug
	case "info":
		return Information
	case "warn":
		return Warning
	case "fatal":
		return Fatal
	case "none":
		return None
	}
	return Information
}
