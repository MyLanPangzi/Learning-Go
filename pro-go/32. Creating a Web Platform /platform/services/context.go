package services

import (
	"context"
	"reflect"
)

type serviceKey string

const (
	Key serviceKey = "key"
)

type serviceMap map[reflect.Type]reflect.Value

func ContextWithService(ctx context.Context) context.Context {
	if ctx.Value(Key) == nil {
		return context.WithValue(ctx, Key, make(serviceMap))
	}
	return ctx
}
