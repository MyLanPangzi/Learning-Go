package logging

import (
	"log"
	"platform/config"
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

func TestNewDefaultLoggerFromConfig(t *testing.T) {
	cfg, err := config.Load("testdata/config.json")
	if err != nil {
		t.Fatal(err)
	}
	logger := NewDefaultLogger(cfg)
	if logger.minLevel != Debug {
		t.Errorf("expected [%v] got [%v]", Debug, logger.minLevel)
	}
}
