package _7

//
//import (
//	"fmt"
//	"reflect"
//)
//
//func convert(src, dest any) (any, bool) {
//	sv := reflect.ValueOf(src)
//	dv := reflect.ValueOf(dest)
//	if sv.Type().ConvertibleTo(dv.Type()) {
//		return sv.Convert(dv.Type()).Interface(), true
//	}
//	return src, false
//}
//func main() {
//	n := "hello"
//	p := 100
//	newVal, ok := convert(p, 100.0)
//	fmt.Printf("%v %v %T\n", ok, newVal, newVal)
//	newVal, ok = convert(n, 100.0)
//	fmt.Printf("%v %v %T\n", ok, newVal, newVal)
//}
