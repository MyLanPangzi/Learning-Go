package main

//
//import (
//	"fmt"
//	"reflect"
//)
//
////	func inspectMethods(s interface{}) {
////	   sType := reflect.TypeOf(s)
////	   if sType.Kind() == reflect.Struct || (sType.Kind() == reflect.Ptr &&
////	           sType.Elem().Kind() == reflect.Struct) {
////	       Printfln("Type: %v, Methods: %v", sType, sType.NumMethod())
////	       for i := 0; i < sType.NumMethod(); i++ {
////	           method := sType.Method(i)
////	           Printfln("Method name: %v, Type: %v",
////	               method.Name, method.Type)
////	       }
////	   }
////	}
//func inspectMethods(s any) {
//	st := reflect.TypeOf(s)
//	if st.Kind() != reflect.Struct && !(st.Kind() == reflect.Pointer && st.Elem().Kind() == reflect.Struct) {
//		return
//	}
//	fmt.Printf("type %v methods %v\n", st, st.NumMethod())
//	for i := 0; i < st.NumMethod(); i++ {
//		m := st.Method(i)
//		fmt.Printf("name %v type %v\n", m.Name, m.Type)
//	}
//}
//func main() {
//	inspectMethods(Purchase{})
//	inspectMethods(&Purchase{})
//}
//
////func main() {
////    inspectMethods(Purchase{})
////    inspectMethods(&Purchase{})
////}
