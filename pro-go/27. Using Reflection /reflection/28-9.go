package main

//
//import (
//	"fmt"
//	"reflect"
//)
//
////	func pickValues(slice interface{}, indices ...int) interface{} {
////	   sliceVal := reflect.ValueOf(slice)
////	   if (sliceVal.Kind() == reflect.Slice) {
////	       newSlice := reflect.MakeSlice(sliceVal.Type(), 0, 10)
////	       for _, index := range indices {
////	           newSlice = reflect.Append(newSlice, sliceVal.Index(index))
////	       }
////	       return newSlice
////	   }
////	   return nil
////	}
//func pickValues(slice any, indices ...int) any {
//	sv := reflect.ValueOf(slice)
//	if sv.Kind() == reflect.Slice {
//		s := reflect.MakeSlice(sv.Type(), 0, 10)
//		for _, index := range indices {
//			s = reflect.Append(s, sv.Index(index))
//		}
//		return s.Interface()
//	}
//	return nil
//}
//func main() {
//	s := []string{"0", "1", "2", "3", "4", "5"}
//	evens := pickValues(s, 0, 2, 4).([]string)
//	for _, v := range evens {
//		fmt.Println(v)
//	}
//}
//
////func main() {
////    name := "Alice"
////    city := "London"
////    hobby := "Running"
////    slice := []string { name, city, hobby, "Bob", "Paris", "Soccer" }
////    picked := pickValues(slice, 0, 3, 5)
////    Printfln("Picked values: %v", picked)
////}
