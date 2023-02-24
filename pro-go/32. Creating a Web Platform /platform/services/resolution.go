package services

import (
	"context"
	"fmt"
	"reflect"
)

func GerService(target any) error {
	return GetServiceForContext(context.Background(), target)
}

func GetServiceForContext(ctx context.Context, target any) error {
	tv := reflect.ValueOf(target)
	if tv.Kind() != reflect.Pointer || !tv.Elem().CanSet() {
		return fmt.Errorf("target cannot set %v", target)
	}
	return resolveServiceFromValue(ctx, tv)
}
