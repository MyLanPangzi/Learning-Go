package services

import (
	"context"
	"fmt"
	"reflect"
)

func Populate(s any) error {
	return PopulateForContext(context.Background(), s)
}

func PopulateForContext(ctx context.Context, s any) error {
	return PopulateForContextWithExtras(ctx, s, map[reflect.Type]reflect.Value{})
}

func PopulateForContextWithExtras(ctx context.Context, s any, m map[reflect.Type]reflect.Value) error {
	sv := reflect.ValueOf(s)
	if sv.Kind() != reflect.Pointer || sv.Elem().Kind() != reflect.Struct || !sv.Elem().CanSet() {
		return fmt.Errorf("type is not a struct %v or cannot set", sv.Type())
	}
	sv = sv.Elem()
	for i := 0; i < sv.NumField(); i++ {
		f := sv.Field(i)
		if f.CanSet() {
			if v, ok := m[f.Type()]; ok {
				f.Set(v)
				continue
			}
			//if f.Kind() == reflect.Pointer {
			//	v := reflect.New(f.Type().Elem())
			//	f.Set(v)
			//	resolveServiceFromValue(ctx, v)
			//	continue
			//}
			resolveServiceFromValue(ctx, f.Addr())
		}

	}

	return nil
}
