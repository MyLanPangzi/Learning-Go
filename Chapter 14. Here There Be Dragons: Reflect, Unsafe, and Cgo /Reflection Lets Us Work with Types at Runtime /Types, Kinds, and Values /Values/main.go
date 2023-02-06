package main

import (
	"fmt"
	"reflect"
)

func main() {

	s := []string{"a", "b", "c"}
	v := reflect.ValueOf(s)
	s2 := v.Interface().([]string)
	fmt.Println(s, s2, v)

	//	i := 10
	//iv := reflect.ValueOf(&i)

	i := 10
	v = reflect.ValueOf(&i)
	fmt.Println(v.Kind(), v.Elem(), v.Type())
	v.Elem().SetInt(20)
	fmt.Println(v.Elem(), i)
	changeInt(&i)
	fmt.Println(i)
	changeIntReflect(&i)
	fmt.Println(i)

}

func changeInt(i *int) {
	*i = 220
}

func changeIntReflect(i *int) {
	v := reflect.ValueOf(i)
	v.Elem().SetInt(230)
}
