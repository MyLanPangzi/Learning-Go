package services

import (
	"context"
	"fmt"
	"reflect"
)

type BindingMap struct {
	factoryFunc reflect.Value
	lifecycle
}

var services = make(map[reflect.Type]BindingMap)

func addService(life lifecycle, factoryFunc any) error {
	ft := reflect.TypeOf(factoryFunc)
	if ft.Kind() != reflect.Func || ft.NumOut() != 1 {
		return fmt.Errorf("type cannot be used as service: %v", ft)
	}
	services[ft.Out(0)] = BindingMap{
		factoryFunc: reflect.ValueOf(factoryFunc),
		lifecycle:   life,
	}
	return nil
}

var contextRef = (*context.Context)(nil)
var contextRefType = reflect.TypeOf(contextRef)

func resolveServiceFromValue(ctx context.Context, val reflect.Value) error {
	st := val.Elem().Type()
	if st == contextRefType {
		val.Elem().Set(reflect.ValueOf(ctx))
		return nil
	}
	if binding, found := services[st]; found {
		if binding.lifecycle == Scoped {
			resolveScopedService(ctx, val, binding)
		} else {
			val.Elem().Set(invokeFunction(ctx, binding.factoryFunc)[0])
		}
		return nil
	}
	return fmt.Errorf("cannot find service %v", st)
}

func resolveScopedService(ctx context.Context, val reflect.Value, binding BindingMap) {
	m, ok := ctx.Value(Key).(serviceMap)
	if ok {
		v, ok := m[val.Type()]
		if !ok {
			v = invokeFunction(ctx, binding.factoryFunc)[0]
			m[val.Type()] = v
		}
		val.Elem().Set(v)
		return
	}
	val.Elem().Set(invokeFunction(ctx, binding.factoryFunc)[0])
}

func invokeFunction(ctx context.Context, factoryFunc reflect.Value, args ...any) []reflect.Value {
	return factoryFunc.Call(resolveFuncArgs(ctx, factoryFunc, args...))
}

func resolveFuncArgs(ctx context.Context, factoryFunc reflect.Value, args ...any) []reflect.Value {
	params := make([]reflect.Value, factoryFunc.Type().NumIn())
	if args != nil {
		for i := 0; i < len(args); i++ {
			params[i] = reflect.ValueOf(args[i])
		}
	}
	for i := len(args); i < len(params); i++ {
		pv := reflect.New(factoryFunc.Type().In(i))
		err := resolveServiceFromValue(ctx, pv)
		if err != nil {
			panic(err)
		}
		params[i] = pv.Elem()
	}
	return params
}
