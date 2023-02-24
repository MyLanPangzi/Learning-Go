package services

import (
	"context"
	"fmt"
	"reflect"
)

func Call(f any, args ...any) ([]any, error) {
	return CallForContext(context.Background(), f, args...)
}

func CallForContext(ctx context.Context, f any, args ...any) ([]any, error) {
	fv := reflect.ValueOf(f)
	if fv.Kind() != reflect.Func {
		return nil, fmt.Errorf("f is not a func %v", fv.Type())
	}
	values := invokeFunction(ctx, fv, args...)
	r := make([]any, len(values))
	for i := range r {
		r[i] = values[i].Interface()
	}
	return r, nil
}
