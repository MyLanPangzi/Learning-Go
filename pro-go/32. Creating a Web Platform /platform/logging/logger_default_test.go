package logging

import (
	"log"
	"strings"
	"testing"
)

func TestDefaultLogger_MinLogLevel(t *testing.T) {
	type fields struct {
		minLevel     LogLevel
		loggers      map[LogLevel]*log.Logger
		triggerPanic bool
	}
	type args struct {
		message string
	}
	s := &strings.Builder{}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{"trace", fields{Trace, map[LogLevel]*log.Logger{Trace: log.New(s, "", 0)}, false}, args{"hello"}},
		{"debug", fields{Debug, map[LogLevel]*log.Logger{Debug: log.New(s, "", 0)}, false}, args{"hello"}},
		{"info", fields{Information, map[LogLevel]*log.Logger{Information: log.New(s, "", 0)}, false}, args{"hello"}},
		{"warn", fields{Warning, map[LogLevel]*log.Logger{Warning: log.New(s, "", 0)}, false}, args{"hello"}},
		{"fatal", fields{Fatal, map[LogLevel]*log.Logger{Fatal: log.New(s, "", 0)}, false}, args{"hello"}},
		{"fatal now", fields{Fatal, map[LogLevel]*log.Logger{Fatal: log.New(s, "", 0)}, true}, args{"hello"}},
		{"none", fields{None, map[LogLevel]*log.Logger{Debug: log.New(s, "", 0)}, false}, args{""}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if message := recover(); message != nil && message != tt.args.message {
					t.Errorf("expected [%v] actual [%v]", tt.args.message, message)
				}
			}()
			s.Reset()
			l := &DefaultLogger{
				minLevel:     tt.fields.minLevel,
				loggers:      tt.fields.loggers,
				triggerPanic: tt.fields.triggerPanic,
			}
			write(l, tt.args.message)
			if strings.TrimSpace(s.String()) != tt.args.message {
				t.Errorf("expected [%v] actual [%v]", tt.args.message, s.String())
			}
		})
	}
}
func write(logger *DefaultLogger, message string) {
	switch logger.minLevel {
	case Trace:
		logger.Trace(message)
	case Debug:
		logger.Debug(message)
	case Information:
		logger.Info(message)
	case Warning:
		logger.Warn(message)
	case Fatal:
		logger.Panic(message)
	}
}
