package main

//
//import (
//	"fmt"
//	"reflect"
//)
//
////	func executeFirstVoidMethod(s interface{}) {
////	   sVal := reflect.ValueOf(s)
////	   for i := 0; i < sVal.NumMethod(); i++ {
////	       method := sVal.Type().Method(i)
////	       if method.Type.NumIn() == 1 {
////	           results := method.Func.Call([]reflect.Value{ sVal })
////	           Printfln("Type: %v, Method: %v, Results: %v",
////	               sVal.Type(), method.Name, results)
////	           break
////	       } else {
////	           Printfln("Skipping method %v %v", method.Name, method.Type.NumIn())
////	       }
////	   }
////	}
//func invoke(s any) {
//	sv := reflect.ValueOf(s)
//	for i := 0; i < sv.NumMethod(); i++ {
//		m := sv.Type().Method(i)
//		if m.Type.NumIn() == 1 {
//			r := m.Func.Call([]reflect.Value{sv})
//			fmt.Printf("type %v method %v results %v\n", sv.Type(), m.Name, r)
//			break
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
