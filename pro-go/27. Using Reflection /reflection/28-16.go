package main

//
//import (
//	"fmt"
//	"reflect"
//)
//
//func inspectStructs(ss ...any) {
//	for _, s := range ss {
//		st := reflect.TypeOf(s)
//		if st.Kind() == reflect.Struct {
//			inspect([]int{}, st)
//		}
//	}
//}
//
//func inspect(baseIndex []int, st reflect.Type) {
//	for i := 0; i < st.NumField(); i++ {
//		fieldIndex := append(baseIndex, i)
//		f := st.Field(i)
//		fmt.Printf("index %v name %v type %v exported %v\n",
//			fieldIndex,
//			f.Name,
//			f.Type,
//			f.PkgPath == "",
//		)
//		if f.Type.Kind() == reflect.Struct {
//			inspect(fieldIndex, st.FieldByIndex(fieldIndex).Type)
//		}
//	}
//}
//
//func main() {
//	inspectStructs(Purchase{})
//}
//
////func inspectStructs(structs ...interface{}) {
////	for _, s := range structs {
////		structType := reflect.TypeOf(s)
////		if structType.Kind() == reflect.Struct {
////			inspectStructType([]int{}, structType)
////		}
////	}
////}
////func inspectStructType(baseIndex []int, structType reflect.Type) {
////	Printfln("--- Struct Type: %v", structType)
////	for i := 0; i < structType.NumField(); i++ {
////		fieldIndex := append(baseIndex, i)
////		field := structType.Field(i)
////		fmt.Printf("Field %v: Name: %v, Type: %v, Exported: %v\n",
////			fieldIndex, field.Name, field.Type, field.PkgPath == "")
////		if field.Type.Kind() == reflect.Struct {
////			field := structType.FieldByIndex(fieldIndex)
////			inspectStructType(fieldIndex, field.Type)
////		}
////	}
////	fmt.Printf("--- End Struct Type: %v\n", structType)
////}
////func main() {
////	inspectStructs(Purchase{})
////}
