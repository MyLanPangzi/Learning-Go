package main

//
//import (
//	"fmt"
//	"reflect"
//	"strings"
//)
//
////	func createMap(slice interface{}, op func(interface{}) interface{}) interface{} {
////	   sliceVal := reflect.ValueOf(slice)
////	   if (sliceVal.Kind() == reflect.Slice) {
////	       mapType := reflect.MapOf(sliceVal.Type().Elem(), sliceVal.Type().Elem())
////	       mapVal := reflect.MakeMap(mapType)
////	       for i := 0; i < sliceVal.Len(); i++ {
////	           elemVal := sliceVal.Index(i)
////	           mapVal.SetMapIndex(elemVal, reflect.ValueOf(op(elemVal.Interface())))
////	       }
////	       return mapVal.Interface()
////	   }
////	   return nil
////	}
//func createMap(s any, op func(any) any) any {
//	sv := reflect.ValueOf(s)
//	if sv.Kind() == reflect.Slice {
//		mt := reflect.MapOf(sv.Type().Elem(), sv.Type().Elem())
//		mv := reflect.MakeMap(mt)
//		for i := 0; i < sv.Len(); i++ {
//			mv.SetMapIndex(sv.Index(i), reflect.ValueOf(op(sv.Index(i).Interface())))
//		}
//		return mv.Interface()
//	}
//	return nil
//}
//func main() {
//	m := createMap([]string{"a", "b", "c"}, func(v any) any {
//		if s, ok := v.(string); ok {
//			return strings.ToUpper(s)
//		}
//		return v
//	}).(map[string]string)
//	for k, v := range m {
//		fmt.Println(k, v)
//	}
//}
//
////func main() {
////    names := []string { "Alice", "Bob", "Charlie"}
////    reverse := func(val interface{}) interface{} {
////        if str, ok := val.(string); ok {
////            return strings.ToUpper(str)
////        }
////        return val
////    }
////    namesMap := createMap(names, reverse).(map[string]string)
////    for k, v := range namesMap {
////        Printfln("Key: %v, Value:%v", k, v)
////    }
////}
