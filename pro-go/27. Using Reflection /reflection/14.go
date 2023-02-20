package main

//
//import (
//	"reflect"
//	"strings"
//)
//
//func incrementOrUpper(values ...any) {
//	for _, e := range values {
//		v := reflect.ValueOf(e)
//		if v.Kind() == reflect.Pointer {
//			v = v.Elem()
//		}
//		if v.CanSet() {
//			switch v.Kind() {
//			case reflect.Int:
//				v.SetInt(v.Int() + 1)
//			case reflect.String:
//				v.SetString(strings.ToUpper(v.String()))
//			}
//			Printfln("Modified value :%v", v)
//			continue
//		}
//		Printfln("cannot set %v value :%v", v.Kind(), v)
//	}
//}
//
//func main() {
//	n := "hello"
//	p := 100
//	c := "london"
//	incrementOrUpper(n, p, c)
//	for _, v := range []any{n, p, c} {
//		Printfln("v:%v", v)
//	}
//
//	incrementOrUpper(&n, &p, &c)
//	for _, v := range []any{n, p, c} {
//		Printfln("v:%v", v)
//	}
//}
