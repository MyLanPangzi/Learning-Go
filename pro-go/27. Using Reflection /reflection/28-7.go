package main

//
//import (
//	"fmt"
//	"reflect"
//)
//
////	func enumerateStrings(arrayOrSlice interface{}) {
////	   arrayOrSliceVal := reflect.ValueOf(arrayOrSlice)
////	   if (arrayOrSliceVal.Kind() == reflect.Array ||
////	           arrayOrSliceVal.Kind() == reflect.Slice) &&
////	           arrayOrSliceVal.Type().Elem().Kind() == reflect.String {
////	       for i := 0; i < arrayOrSliceVal.Len(); i++ {
////	           Printfln("Element: %v, Value: %v", i, arrayOrSliceVal.Index(i).String())
////	       }
////	   }
////	}
//func enumerateStrings(arr any) {
//	av := reflect.ValueOf(arr)
//	if (av.Kind() == reflect.Array || av.Kind() == reflect.Slice) && av.Type().Elem().Kind() == reflect.String {
//		for i := 0; i < av.Len(); i++ {
//			fmt.Printf("%v \n", av.Index(i).String())
//		}
//	}
//}
//func main() {
//	enumerateStrings([]string{"1", "2"})
//	enumerateStrings([...]string{"3", "4"})
//}
//
////func main() {
////    name := "Alice"
////    city := "London"
////    hobby := "Running"
////    slice := []string { name, city, hobby }
////    array := [3]string { name, city, hobby}
////    enumerateStrings(slice)
////    enumerateStrings(array)
////}
