package main

//
//import (
//	"fmt"
//	"reflect"
//)
//
//func inspectStructs(structs ...any) {
//	for _, s := range structs {
//		st := reflect.TypeOf(s)
//		if st.Kind() == reflect.Struct {
//			inspectStruct(st)
//		}
//	}
//}
//
//func inspectStruct(st reflect.Type) {
//	for i := 0; i < st.NumField(); i++ {
//		f := st.Field(i)
//		fmt.Printf("index %v name %v type %v exported %v \n",
//			f.Index,
//			f.Name,
//			f.Type,
//			f.PkgPath == "",
//		)
//	}
//}
//func main() {
//	inspectStructs(Purchase{})
//}
//
////func inspectStructs(structs ...interface{}) {
////    for _, s := range structs {
////        structType := reflect.TypeOf(s)
////        if (structType.Kind() == reflect.Struct) {
////            inspectStructType(structType)
////        }
////    }
////}
////func inspectStructType(structType reflect.Type) {
////    Printfln("--- Struct Type: %v", structType)
////    for i := 0; i < structType.NumField(); i++ {
////        field := structType.Field(i)
////        Printfln("Field %v: Name: %v, Type: %v, Exported: %v",
////            field.Index, field.Name, field.Type, field.PkgPath == "")
////    }
////    Printfln("--- End Struct Type: %v", structType)
////}
////func main() {
////    inspectStructs( Purchase{} )
////}
