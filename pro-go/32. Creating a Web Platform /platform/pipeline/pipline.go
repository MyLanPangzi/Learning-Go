package pipeline

import (
	"fmt"
	"net/http"
	"platform/services"
	"reflect"
)

type RequestPipeline func(*ComponentContext)

var emptyPipeline RequestPipeline = func(*ComponentContext) {}

func CreatePipeline(components ...any) RequestPipeline {
	f := emptyPipeline
	for i := len(components) - 1; i >= 0; i-- {
		current := components[i]
		err := services.Populate(current)
		if err != nil {
			panic(err)
		}
		next := f
		if svcComp, ok := current.(ServiceMiddlewareComponent); ok {
			f = createServiceDependentFunc(current, next)
			svcComp.Init()
		} else if stdComp, ok := current.(MiddlewareComponent); ok {
			f = func(context *ComponentContext) {
				if context.error == nil {
					stdComp.ProcessRequest(context, next)
				}
			}
			stdComp.Init()
		} else {
			panic(fmt.Sprintf("%v is not middleware component %T", current, current))
		}
	}
	return f
}

func createServiceDependentFunc(current any, next RequestPipeline) RequestPipeline {
	method := reflect.ValueOf(current).MethodByName("ProcessRequestWithServices")
	if !method.IsValid() {
		panic(fmt.Sprintf("%v  did not implements ProcessRequestWithServices %T", current, current))
	}
	return func(context *ComponentContext) {
		if context.error == nil {
			_, err := services.CallForContext(context.Context(), method.Interface(), context, next)
			if err != nil {
				context.Error(err)
			}
		}
	}
}

//	func CreatePipeline(components ...MiddlewareComponent) RequestPipeline {
//		f := emptyPipeline
//		for i := len(components) - 1; i >= 0; i-- {
//			current := components[i]
//			next := f
//			f = func(ctx *ComponentContext) {
//				if ctx.error == nil {
//					current.ProcessRequest(ctx, next)
//				}
//			}
//			current.Init()
//		}
//		return f
//	}
func (pl RequestPipeline) ProcessRequest(r *http.Request, w http.ResponseWriter) error {
	ctx := ComponentContext{r, w, nil}
	pl(&ctx)
	return ctx.error
}
