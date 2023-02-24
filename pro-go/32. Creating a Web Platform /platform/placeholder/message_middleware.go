package placeholder

import (
	"errors"
	"io"
	"platform/config"
	"platform/pipeline"
	"platform/services"
)

type SimpleMessageComponent struct {
}

func (s *SimpleMessageComponent) Init() {
}

func (s *SimpleMessageComponent) ProcessRequest(context *pipeline.ComponentContext, next func(*pipeline.ComponentContext)) {
	var cfg config.Configuration
	err := services.GetService(&cfg)
	if err != nil {
		panic(err)
	}
	msg, found := cfg.GetString("main:message")
	if found {
		_, err := io.WriteString(context.ResponseWriter, msg)
		if err != nil {
			panic(err)
		}
	} else {
		context.Error(errors.New("cannot find config setting"))
	}
	//next(context)
}
