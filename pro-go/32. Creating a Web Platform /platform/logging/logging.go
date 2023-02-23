package logging

type LogLevel int

const (
	Trace LogLevel = iota
	Debug
	Information
	Warning
	Fatal
	None
)

type Logger interface {
	Trace(message string)
	Tracef(template string, args ...any)
	Debug(message string)
	Debugf(template string, args ...any)
	Info(message string)
	Infof(template string, args ...any)
	Warn(message string)
	Warnf(template string, args ...any)
	Panic(message string)
	Panicf(template string, args ...any)
}
