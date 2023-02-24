package basic

import (
	"fmt"
	"net/http"
	"platform/logging"
	"platform/pipeline"
	"platform/services"
)

type ErrorComponent struct {
}

func (e *ErrorComponent) Init() {
}

func (e *ErrorComponent) ProcessRequest(
	context *pipeline.ComponentContext,
	next func(*pipeline.ComponentContext),
) {
	var logger logging.Logger
	err := services.GetServiceForContext(context.Context(), &logger)
	if err != nil {
		panic(err)
	}
	defer func() {
		if arg := recover(); arg != nil {
			logger.Debugf("Error:%v", fmt.Sprint(arg))
			context.WriteHeader(http.StatusInternalServerError)
		}
	}()
	next(context)
	if context.GetError() != nil {
		logger.Debugf("Error:%v", context.GetError())
		context.WriteHeader(http.StatusInternalServerError)
	}
}
