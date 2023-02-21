package main

//
//import (
//	"fmt"
//	"reflect"
//)
//
////func main() {
////    name := "Alice"
////    city := "London"
////    hobby := "Running"
////    slice := []string { name, city, hobby }
////    array := [3]string { name, city, hobby}
////    Printfln("Slice (string): %v",  checkElemType("testString", slice))
////    Printfln("Array (string): %v",  checkElemType("testString", array))
////    Printfln("Array (int): %v",  checkElemType(10, array))
////}
//
//func checkElemType(val, arr any) bool {
//	vt := reflect.TypeOf(val)
//	at := reflect.TypeOf(arr)
//	return (at.Kind() == reflect.Array || at.Kind() == reflect.Slice) && at.Elem() == vt
//}
//func main() {
//	at := reflect.ArrayOf(3, reflect.TypeOf("s"))
//	st := reflect.SliceOf(reflect.TypeOf("s"))
//	fmt.Println(at == reflect.TypeOf([3]string{}), st == reflect.TypeOf(([]string)(nil)))
//	n := "hello"
//	c := "hu"
//	h := "run"
//	s := []string{n, c, h}
//	a := [...]string{n, c, h}
//	fmt.Println(checkElemType(n, s), checkElemType(c, a), checkElemType(10, a))
//}
