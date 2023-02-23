package main

//
//import (
//	"fmt"
//	"reflect"
//	"strings"
//)
//
//func Map(s, f any) []any {
//	sv := reflect.ValueOf(s)
//	fv := reflect.ValueOf(f)
//	if sv.Kind() != reflect.Slice || fv.Kind() != reflect.Func {
//		return nil
//	}
//	inTypes := []reflect.Type{sv.Type().Elem()}
//	var outTypes []reflect.Type
//	for i := 0; i < fv.Type().NumOut(); i++ {
//		outTypes = append(outTypes, fv.Type().Out(i))
//	}
//	ft := reflect.FuncOf(inTypes, outTypes, fv.Type().IsVariadic())
//	if ft != fv.Type() {
//		return nil
//	}
//	var out []any
//	for i := 0; i < sv.Len(); i++ {
//		out = append(out, fv.Call([]reflect.Value{sv.Index(i)})[0].Interface())
//	}
//	return out
//}
//func main() {
//	fmt.Println(Map([]string{"a", "b"}, strings.ToUpper))
//}
//
////func mapSlice(slice interface{}, mapper interface{}) (mapped []interface{}) {
////    sliceVal := reflect.ValueOf(slice)
////    mapperVal := reflect.ValueOf(mapper)
////    mapped = []interface{} {}
////    if sliceVal.Kind() == reflect.Slice && mapperVal.Kind() == reflect.Func {
////        paramTypes := []reflect.Type { sliceVal.Type().Elem() }
////        resultTypes := []reflect.Type {}
////        for i := 0; i < mapperVal.Type().NumOut(); i++ {
////            resultTypes = append(resultTypes, mapperVal.Type().Out(i))
////        }
////        expectedFuncType := reflect.FuncOf(paramTypes,
////            resultTypes, mapperVal.Type().IsVariadic())
////        if (mapperVal.Type() == expectedFuncType) {
////            for i := 0; i < sliceVal.Len(); i++ {
////                result := mapperVal.Call([]reflect.Value {sliceVal.Index(i)})
////                for _, r := range result {
////                    mapped = append(mapped, r.Interface())
////                }
////            }
////        } else {
////            Printfln("Function type not as expected")
////        }
////    }
////    return
////}
////func main() {
////    names := []string { "Alice", "Bob", "Charlie" }
////    results := mapSlice(names, strings.ToUpper)
////    Printfln("Results: %v", results)
////}
