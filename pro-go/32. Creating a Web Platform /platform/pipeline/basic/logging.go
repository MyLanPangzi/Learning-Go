package basic

import (
	"net/http"
	"platform/logging"
	"platform/pipeline"
)

type LoggingResponseWriter struct {
	statusCode int
	http.ResponseWriter
}

func (l *LoggingResponseWriter) WriteHeader(statusCode int) {
	l.statusCode = statusCode
	l.ResponseWriter.WriteHeader(statusCode)
}
func (l *LoggingResponseWriter) Write(b []byte) (int, error) {
	if l.statusCode == 0 {
		l.statusCode = http.StatusOK
	}
	return l.ResponseWriter.Write(b)
}

type LoggingComponent struct {
}

func (l *LoggingComponent) ImplementsProcessRequestWithServices() {
}

func (l *LoggingComponent) Init() {
}

func (l *LoggingComponent) ProcessRequestWithServices(
	context *pipeline.ComponentContext,
	next func(*pipeline.ComponentContext),
	logger logging.Logger,
) {
	w := &LoggingResponseWriter{0, context.ResponseWriter}
	context.ResponseWriter = w
	logger.Infof("REQ --- %v - %v", context.Request.Method, context.URL)
	next(context)
	logger.Infof("RESP  %v  %v", w.statusCode, context.URL)
}
