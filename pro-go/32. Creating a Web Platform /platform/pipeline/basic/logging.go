package basic

import (
	"net/http"
	"platform/logging"
	"platform/pipeline"
	"platform/services"
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

func (l *LoggingComponent) Init() {
}

func (l *LoggingComponent) ProcessRequest(
	context *pipeline.ComponentContext,
	next func(*pipeline.ComponentContext),
) {
	var logger logging.Logger
	err := services.GetServiceForContext(context.Context(), &logger)
	if err != nil {
		context.Error(err)
		return
	}
	w := &LoggingResponseWriter{0, context.ResponseWriter}
	context.ResponseWriter = w
	logger.Infof("REQ --- %v - %v", context.Request.Method, context.URL)
	logger.Infof("RESP  %v  %v", w.statusCode, context.URL)
}

//type LoggingComponent struct {}
//func (lc *LoggingComponent) Init() {}
//func (lc *LoggingComponent) ProcessRequest(ctx *pipeline.ComponentContext,
//        next func(*pipeline.ComponentContext))  {
//    var logger logging.Logger
//    err := services.GetServiceForContext(ctx.Request.Context(), &logger)
//    if (err != nil) {
//        ctx.Error(err)
//        return
//    }
//    loggingWriter := LoggingResponseWriter{ 0, ctx.ResponseWriter}
//    ctx.ResponseWriter = &loggingWriter
//    logger.Infof("REQ --- %v - %v", ctx.Request.Method, ctx.Request.URL)
//    next(ctx)
//    logger.Infof("RSP %v %v", loggingWriter.statusCode, ctx.Request.URL )
//}
