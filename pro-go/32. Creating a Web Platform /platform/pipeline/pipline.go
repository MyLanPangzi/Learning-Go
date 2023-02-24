package pipeline

import "net/http"

type RequestPipeline func(*ComponentContext)

var emptyPipeline RequestPipeline = func(*ComponentContext) {}

func CreatePipeline(components ...MiddlewareComponent) RequestPipeline {
	f := emptyPipeline
	for i := len(components) - 1; i >= 0; i-- {
		current := components[i]
		next := f
		f = func(ctx *ComponentContext) {
			if ctx.error == nil {
				current.ProcessRequest(ctx, next)
			}
		}
		current.Init()
	}
	return f
}
func (pl RequestPipeline) ProcessRequest(r *http.Request, w http.ResponseWriter) error {
	ctx := ComponentContext{r, w, nil}
	pl(&ctx)
	return ctx.error
}
