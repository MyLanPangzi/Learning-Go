package main

//
//import (
//	"fmt"
//	"math"
//	"reflect"
//)
//
//func IsInt(v reflect.Value) bool {
//	switch v.Kind() {
//	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
//		return true
//	}
//	return false
//}
//func IsFloat(v reflect.Value) bool {
//	switch v.Kind() {
//	case reflect.Float32, reflect.Float64:
//		return true
//	}
//	return false
//}
//
//func convert(src, target any) (any, bool) {
//	sv := reflect.ValueOf(src)
//	tv := reflect.ValueOf(target)
//	if sv.Type().ConvertibleTo(tv.Type()) {
//		switch {
//		case IsInt(sv) && IsInt(tv) && tv.OverflowInt(sv.Int()):
//			fmt.Println("int overflow")
//			return src, false
//		case IsFloat(sv) && IsFloat(tv) && tv.OverflowFloat(sv.Float()):
//			fmt.Println("float overflow")
//			return src, false
//		default:
//			return sv.Convert(tv.Type()).Interface(), true
//		}
//	}
//	return src, false
//}
//func main() {
//	p := 100
//	n := "hello"
//	v, ok := convert(p, 200.0)
//	fmt.Printf("%v %v %T\n", ok, v, v)
//	v, ok = convert(n, 200.0)
//	fmt.Printf("%v %v %T\n", ok, v, v)
//	v, ok = convert(5000, int8(100))
//	fmt.Printf("%v %v %T\n", ok, v, v)
//	v, ok = convert(math.MaxFloat64, float32(math.MaxFloat32))
//	fmt.Printf("%v %v %T\n", ok, v, v)
//}
//
////func main() {
////    name := "Alice"
////    price := 279
////    //city := "London"
////    newVal, ok := convert(price, 100.00)
////    Printfln("Converted %v: %v, %T", ok, newVal, newVal)
////    newVal, ok = convert(name, 100.00)
////    Printfln("Converted %v: %v, %T", ok, newVal, newVal)
////    newVal, ok = convert(5000, int8(100))
////    Printfln("Converted %v: %v, %T", ok, newVal, newVal)
////}
