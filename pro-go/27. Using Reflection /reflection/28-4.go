package main

//
//import (
//	"fmt"
//	"reflect"
//	"strings"
//)
//
//var (
//	stringPtrType = reflect.TypeOf((*string)(nil))
//)
//
//func transformString(val any) {
//	v := reflect.ValueOf(val)
//	if v.Type() == stringPtrType {
//		upper := strings.ToUpper(v.Elem().String())
//		if v.Elem().CanSet() {
//			v.Elem().SetString(upper)
//		}
//	}
//}
//func main() {
//	n := "hello"
//	transformString(&n)
//	fmt.Println(n)
//}
