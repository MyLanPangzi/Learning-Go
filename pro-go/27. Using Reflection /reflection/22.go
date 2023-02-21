package main

//
//import (
//	"fmt"
//	"reflect"
//)
//
//func swap(first, second any) {
//	fv, sv := reflect.ValueOf(first), reflect.ValueOf(second)
//	if fv.Type() == sv.Type() &&
//		fv.Kind() == reflect.Pointer &&
//		sv.Kind() == reflect.Pointer &&
//		fv.Elem().CanSet() &&
//		sv.Elem().CanSet() {
//		v := reflect.New(fv.Elem().Type())
//		v.Elem().Set(fv.Elem())
//		fv.Elem().Set(sv.Elem())
//		sv.Elem().Set(v.Elem())
//	}
//}
//
//func swap2[T any](x, y *T) {
//	*x, *y = *y, *x
//}
//
//
//func main() {
//	n := "hello"
//	//p := 100
//	c := "hu"
//	swap(&n, &c)
//	fmt.Println(n, c)
//	swap2(&n, &c)
//	fmt.Println(n, c)
//
//}
