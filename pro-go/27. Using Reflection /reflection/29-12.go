package main

//
//import (
//	"fmt"
//	"reflect"
//)
//
////func executeFirstVoidMethod(s interface{}) {
////    sVal := reflect.ValueOf(s)
////    for i := 0; i < sVal.NumMethod(); i++ {
////        method := sVal.Method(i)
////        if method.Type().NumIn() == 0 {
////            results := method.Call([]reflect.Value{})
////            Printfln("Type: %v, Method: %v, Results: %v",
////                sVal.Type(), sVal.Type().Method(i).Name, results)
////            break
////        } else {
////            Printfln("Skipping method %v %v",
////                sVal.Type().Method(i).Name, method.Type().NumIn())
////        }
////    }
////}
//
//func invoke(s any) {
//	sv := reflect.ValueOf(s)
//	for i := 0; i < sv.NumMethod(); i++ {
//		m := sv.Method(i)
//		if m.Type().NumIn() == 0 {
//			r := m.Call([]reflect.Value{})
//			fmt.Printf(
//				"type %v method %v result %v \n",
//				sv.Type(),
//				sv.Type().Method(i).Name,
//				r,
//			)
//		}
//	}
//}
//func main() {
//	invoke(&Product{"apple", "fruit", 1})
//}
//
////func main() {
////    executeFirstVoidMethod(&Product { Name: "Kayak", Price: 279})
////}
