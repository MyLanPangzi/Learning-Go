package main

//
//import (
//	"fmt"
//	"reflect"
//)
//
////	func getUnderlying(item Wrapper, fieldName string) {
////	   itemVal := reflect.ValueOf(item)
////	   fieldVal := itemVal.FieldByName(fieldName)
////	   Printfln("Field Type: %v", fieldVal.Type())
////	   for i := 0; i < fieldVal.Type().NumMethod(); i++ {
////	       method := fieldVal.Type().Method(i)
////	       Printfln("Interface Method: %v, Exported: %v",
////	           method.Name, method.PkgPath == "")
////	   }
////	   Printfln("--------")
////	   if (fieldVal.Kind() == reflect.Interface) {
////	       Printfln("Underlying Type: %v", fieldVal.Elem().Type())
////	       for i := 0; i < fieldVal.Elem().Type().NumMethod(); i++ {
////	           method := fieldVal.Elem().Type().Method(i)
////	           Printfln("Underlying Method: %v", method.Name)
////	       }
////	   }
////	}
//type Wrapper struct {
//	NamedItem
//}
//
//func getUnderlying(item Wrapper, name string) {
//	iv := reflect.ValueOf(item)
//	f := iv.FieldByName(name)
//	fmt.Printf("type %v\n", f.Type())
//	for i := 0; i < f.Type().NumMethod(); i++ {
//		m := f.Type().Method(i)
//		m.IsExported()
//		fmt.Printf("interface method %v, Exported %v\n", m.Name, m.PkgPath == "")
//	}
//	fmt.Println()
//	if f.Kind() == reflect.Interface {
//		fmt.Printf("underlying type %v\n", f.Elem().Type())
//		for i := 0; i < f.Elem().Type().NumMethod(); i++ {
//			m := f.Elem().Type().Method(i)
//			fmt.Printf("underlying type method %v, Exported %v\n", m.Name, m.PkgPath == "")
//		}
//	}
//}
//func main() {
//	getUnderlying(Wrapper{NamedItem: &Product{}}, "NamedItem")
//}
