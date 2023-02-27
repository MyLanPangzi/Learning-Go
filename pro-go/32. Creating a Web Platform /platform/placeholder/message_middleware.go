package placeholder

import (
	"platform/config"
	"platform/pipeline"
	"platform/templates"
)

type SimpleMessageComponent struct {
	Message string
	Cfg     config.Configuration
}

func (s *SimpleMessageComponent) ImplementsProcessRequestWithServices() {
}

func (s *SimpleMessageComponent) Init() {
	s.Message = s.Cfg.GetStringDefault("main:message", "Hello")
}

func (s *SimpleMessageComponent) ProcessRequestWithServices(
	context *pipeline.ComponentContext,
	next func(*pipeline.ComponentContext),
	t templates.TemplateExecutor,
) {
	err := t.ExecTemplate(context.ResponseWriter, "simple_message.go.html", s.Message)
	if err != nil {
		context.Error(err)
		return
	}
	next(context)
}
