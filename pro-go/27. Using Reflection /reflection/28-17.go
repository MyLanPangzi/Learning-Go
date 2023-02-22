package main

//
//import (
//	"fmt"
//	"reflect"
//)
//
//func descField(s any, fieldName string) {
//	st := reflect.TypeOf(s)
//	if f, ok := st.FieldByName(fieldName); ok {
//		fmt.Printf("name %v type %v index %v\n", f.Name, f.Type, f.Index)
//		for index := f.Index; len(index) > 1; {
//			index = index[0 : len(index)-1]
//			f := st.FieldByIndex(index)
//			fmt.Printf("parent %v type %v index %v \n",
//				f.Name,
//				f.Type,
//				f.Index,
//			)
//		}
//		fmt.Printf("top level name %v type %v \n", st.Name(), st)
//	}
//
//}
//func main() {
//	descField(Purchase{}, "Price")
//}
//
////func describeField(s interface{}, fieldName string) {
////    structType := reflect.TypeOf(s)
////    field, found := structType.FieldByName(fieldName)
////    if (found) {
////        Printfln("Found: %v, Type: %v, Index: %v",
////            field.Name, field.Type, field.Index)
////        index := field.Index
////        for len(index) > 1 {
////            index = index[0: len(index) -1]
////            field = structType.FieldByIndex(index)
////            Printfln("Parent : %v, Type: %v, Index: %v",
////                field.Name, field.Type, field.Index)
////        }
////        Printfln("Top-Level Type: %v" , structType)
////    } else {
////        Printfln("Field %v not found", fieldName)
////    }
////}
////func main() {
////    describeField( Purchase{}, "Price" )
////}
