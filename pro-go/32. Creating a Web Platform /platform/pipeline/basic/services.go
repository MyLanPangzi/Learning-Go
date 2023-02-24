package basic

import (
	"platform/pipeline"
	"platform/services"
)

type ServicesComponent struct {
}

func (s ServicesComponent) Init() {
}

func (s ServicesComponent) ProcessRequest(
	ctx *pipeline.ComponentContext,
	next func(*pipeline.ComponentContext),
) {
	context := ctx.Request.Context()
	r := ctx.Request.WithContext(services.ContextWithService(context))
	ctx.Request = r
	next(ctx)
}
