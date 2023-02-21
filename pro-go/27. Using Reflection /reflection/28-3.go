package main

//
//import (
//	"fmt"
//	"reflect"
//)
//
//func createPointerType(t reflect.Type) reflect.Type {
//	return reflect.PointerTo(t)
//}
//func followPointerType(t reflect.Type) reflect.Type {
//	if t.Kind() == reflect.Pointer {
//		return t.Elem()
//	}
//	return t
//}
//func main() {
//	n := "hello"
//	t := reflect.TypeOf(n)
//	pt := createPointerType(t)
//	fmt.Printf("%v %v %v %v\n", pt, followPointerType(pt), t, followPointerType(pt) == t)
//
//}
