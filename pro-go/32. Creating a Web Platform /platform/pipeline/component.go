package pipeline

import "net/http"

type ComponentContext struct {
	*http.Request
	http.ResponseWriter
	error
}

func (c *ComponentContext) Error(err error) {
	c.error = err
}
func (c *ComponentContext) GetError() error {
	return c.error
}

type MiddlewareComponent interface {
	Init()
	ProcessRequest(*ComponentContext, func(*ComponentContext))
}

type ServiceMiddlewareComponent interface {
	Init()
	ImplementsProcessRequestWithServices()
}
