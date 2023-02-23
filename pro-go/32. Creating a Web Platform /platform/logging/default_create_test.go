package logging

import (
	"log"
	"strings"
	"testing"
)

func TestDefaultLoggerPrefix(t *testing.T) {
	t.Run("TRACE", func(t *testing.T) {
		r := &strings.Builder{}
		l := &DefaultLogger{
			minLevel: Trace,
			loggers: map[LogLevel]*log.Logger{
				Trace: log.New(r, "TRACE ", log.Lmsgprefix),
			},
			triggerPanic: false,
		}
		l.Trace("hello")
		if strings.TrimSpace(r.String()) != "TRACE hello" {
			t.Errorf("expected [%v] got [%v]", "TRACE hello", strings.TrimSpace(r.String()))
		}
	})
}
