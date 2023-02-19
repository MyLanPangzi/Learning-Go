package main

import (
	"fmt"
	"reflect"
)

func main() {
	//  var names [3]string
	//    names[0] = "Kayak"
	//    names[1] = "Lifejacket"
	//    names[2] = "Paddle"
	//    fmt.Println(names)
	var names [3]string
	fmt.Println(names)
	names[0] = "hello"
	names[1] = "world"
	names[2] = "go"
	fmt.Printf("%T %v\n", names, names)
	t := reflect.TypeOf(names)
	fmt.Println(t, t.Len(), t.Elem())
	v := reflect.ValueOf(names)
	fmt.Println(v.Kind(), v.Type(), v.Len())
}
