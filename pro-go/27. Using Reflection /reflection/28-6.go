package main

//
//import (
//	"fmt"
//	"reflect"
//)
//
//func setValue(arr any, i int, r any) {
//	av := reflect.ValueOf(arr)
//	rv := reflect.ValueOf(r)
//	if av.Kind() == reflect.Slice {
//		if v := av.Index(i); v.CanSet() {
//			v.Set(rv)
//		}
//		return
//	}
//	if av.Kind() == reflect.Pointer &&
//		av.Elem().Kind() == reflect.Array &&
//		av.Elem().Type().Elem() == rv.Type() &&
//		av.Elem().CanSet() {
//		av.Elem().Index(i).Set(rv)
//	}
//
//}
//func main() {
//	n := "hello"
//	c := "lon"
//	s := []string{n, c}
//	a := [...]string{n, c}
//	fmt.Println(a, s)
//	setValue(s, 1, "hu")
//	fmt.Println(a, s)
//	setValue(&a, 1, "nan")
//	fmt.Println(a, s)
//	fmt.Println(reflect.ValueOf(&a).Elem().Type(), reflect.ValueOf(&a).Elem().Type().Elem())
//	av := reflect.ValueOf(a)
//	sv := reflect.ValueOf(s)
//	fmt.Println(av.Len(), av.Cap(), sv.Len(), sv.Cap())
//	newS := sv.Slice(0, 1)
//	fmt.Println(newS.Len(), newS.Cap())
//	newS = sv.Slice3(0, 1, 1)
//	fmt.Println(newS.Len(), newS.Cap())
//	newS = sv.Slice3(0, 1, 2)
//	fmt.Println(newS.Len(), newS.Cap())
//	//sv.SetCap(10)
//	//sv.SetLen(5)
//	//fmt.Println(sv.Len(), sv.Cap())
//
//}
//
////func main() {
////    name := "Alice"
////    city := "London"
////    hobby := "Running"
////    slice := []string { name, city, hobby }
////    array := [3]string { name, city, hobby}
////    Printfln("Original slice: %v", slice)
////    newCity := "Paris"
////    setValue(slice, 1, newCity)
////    Printfln("Modified slice: %v", slice)
////    Printfln("Original slice: %v", array)
////    newCity = "Rome"
////    setValue(&array, 1, newCity)
////    Printfln("Modified slice: %v", array)
////}
