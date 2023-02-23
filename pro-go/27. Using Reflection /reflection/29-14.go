package main

//
//import (
//	"fmt"
//	"reflect"
//)
//
////	type Wrapper struct {
////	   NamedItem
////	}
////
////	func getUnderlying(item Wrapper, fieldName string) {
////	   itemVal := reflect.ValueOf(item)
////	   fieldVal := itemVal.FieldByName(fieldName)
////	   Printfln("Field Type: %v", fieldVal.Type())
////	   if (fieldVal.Kind() == reflect.Interface) {
////	       Printfln("Underlying Type: %v", fieldVal.Elem().Type())
////	   }
////	}
////
////	func main() {
////	   getUnderlying(Wrapper{NamedItem: &Product{}}, "NamedItem")
////	}
//type Wrapper struct {
//	NamedItem
//}
//
//func getUnderlying(item Wrapper, fieldName string) {
//	iv := reflect.ValueOf(item)
//	f := iv.FieldByName(fieldName)
//	fmt.Printf("type %v\n", f.Type())
//	if f.Kind() == reflect.Interface {
//		fmt.Printf("underlying type %v\n", f.Elem().Type())
//
//	}
//}
//func main() {
//	getUnderlying(Wrapper{&Product{}}, "NamedItem")
//}
