package services

import (
	"fmt"
	"reflect"
	"sync"
)

func AddTransient(f any) error {
	return addService(Transient, f)
}

func AddScoped(f any) error {
	return addService(Scoped, f)
}
func AddSingleton(f any) error {
	fv := reflect.ValueOf(f)
	if fv.Kind() != reflect.Func || fv.Type().NumOut() != 1 {
		return fmt.Errorf("func is not service factory func :%v", fv.Type())
	}
	var results []reflect.Value
	once := sync.Once{}
	wrapper := reflect.MakeFunc(fv.Type(), func(args []reflect.Value) []reflect.Value {
		once.Do(func() {
			results = invokeFunction(nil, fv)
		})
		return results
	}).Interface()
	return addService(Singleton, wrapper)
}
