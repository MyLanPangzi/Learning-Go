package main

import (
	"fmt"
	"reflect"
)

func main() {
	stringType := reflect.TypeOf((*string)(nil)).Elem()
	stringSliceType := reflect.TypeOf([]string(nil))
	v := reflect.New(stringType).Elem()
	v.SetString("hello")
	fmt.Println(v)

	slice := reflect.MakeSlice(stringSliceType, 0, 10)
	slice = reflect.Append(slice, v)
	s := slice.Interface().([]string)
	fmt.Println(s)

	intSliceType := reflect.SliceOf(reflect.TypeOf(10))
	intSlice := reflect.MakeSlice(intSliceType, 0, 10)
	intSlice = reflect.Append(intSlice, reflect.ValueOf(10))
	fmt.Println(intSlice)
}
