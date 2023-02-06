package main

import (
	"fmt"
	"reflect"
)

type Foo struct {
	A int    `myTag:"value"`
	B string `myTag:"value2"`
}

var x Foo

func DoSomething(f Foo) {
	fmt.Println(f.A, f.B)
}
func main() {
	var x int
	t := reflect.TypeOf(x)
	fmt.Println(t.Name())
	fmt.Println(t.Kind())
	f := Foo{
		A: 0,
		B: "",
	}
	t = reflect.TypeOf(f)
	fmt.Println(t.Name())
	fmt.Println(t.Kind())

	p := &x
	t = reflect.TypeOf(p)
	fmt.Println(t.Name())
	fmt.Println(t.Kind(), t.Kind() == reflect.Ptr, t.Kind() == reflect.Pointer)
	fmt.Println(t.Elem().Name())
	fmt.Println(t.Elem().Kind(), t.Elem().Kind() == reflect.Int)

	t = reflect.TypeOf(f)
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		fmt.Println(f.Name, f.Tag, f.Type.Name(), f.Index, f.PkgPath)
	}
}
