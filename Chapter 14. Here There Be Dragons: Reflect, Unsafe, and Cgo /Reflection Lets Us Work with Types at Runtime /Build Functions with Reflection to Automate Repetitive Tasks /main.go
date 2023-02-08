package main

import (
	"errors"
	"fmt"
	"reflect"
	"time"
)

func main() {
	f := MakeTimedFunction(timeMe).(func(int) int)
	fmt.Println(f(2))
}

func MakeTimedFunction(f any) any {
	ft := reflect.TypeOf(f)
	fv := reflect.ValueOf(f)
	if fv.Kind() != reflect.Func {
		panic(errors.New("expects a func"))
	}
	return reflect.MakeFunc(ft, func(in []reflect.Value) []reflect.Value {
		now := time.Now()
		out := fv.Call(in)
		fmt.Println(time.Since(now))
		return out
	}).Interface()
}

func timeMe(a int) int {
	time.Sleep(time.Duration(a) * time.Second)
	return a * 2
}
