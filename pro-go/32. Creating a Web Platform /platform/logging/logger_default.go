package logging

import (
	"fmt"
	"io"
	"log"
)

type DefaultLogger struct {
	minLevel     LogLevel
	loggers      map[LogLevel]*log.Logger
	triggerPanic bool
}

func (l *DefaultLogger) MinLogLevel() LogLevel {
	return l.minLevel
}

func (l *DefaultLogger) setLoggerOut(level LogLevel, w io.Writer) {
	l.loggers[level].SetOutput(w)
}
func (l *DefaultLogger) write(level LogLevel, message string) {
	if l.minLevel <= level {
		l.loggers[level].Output(2, message)
	}
}
func (l *DefaultLogger) Trace(message string) {
	l.write(Trace, message)
}

func (l *DefaultLogger) Tracef(template string, args ...any) {
	l.write(Trace, fmt.Sprintf(template, args...))
}

func (l *DefaultLogger) Debug(message string) {
	l.write(Debug, message)
}

func (l *DefaultLogger) Debugf(template string, args ...any) {
	l.write(Debug, fmt.Sprintf(template, args...))
}

func (l *DefaultLogger) Info(message string) {
	l.write(Information, message)
}

func (l *DefaultLogger) Infof(template string, args ...any) {
	l.write(Information, fmt.Sprintf(template, args...))
}

func (l *DefaultLogger) Warn(message string) {
	l.write(Warning, message)
}

func (l *DefaultLogger) Warnf(template string, args ...any) {
	l.write(Warning, fmt.Sprintf(template, args...))
}

func (l *DefaultLogger) Panic(message string) {
	l.write(Fatal, message)
	if l.triggerPanic {
		panic(message)
	}
}

func (l *DefaultLogger) Panicf(template string, args ...any) {
	message := fmt.Sprintf(template, args...)
	l.write(Fatal, message)
	if l.triggerPanic {
		panic(message)
	}
}
