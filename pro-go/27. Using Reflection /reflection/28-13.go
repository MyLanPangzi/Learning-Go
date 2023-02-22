package main

//
//import (
//	"fmt"
//	"reflect"
//)
//
////	func setMap(m interface{}, key interface{}, val interface{}) {
////	   mapValue := reflect.ValueOf(m)
////	   keyValue := reflect.ValueOf(key)
////	   valValue := reflect.ValueOf(val)
////	   if (mapValue.Kind() == reflect.Map &&
////	       mapValue.Type().Key() == keyValue.Type() &&
////	       mapValue.Type().Elem() == valValue.Type()) {
////	           mapValue.SetMapIndex(keyValue, valValue)
////	   } else {
////	       Printfln("Not a map or mismatched types")
////	   }
////	}
//func setMap(m, k, v any) {
//	mv := reflect.ValueOf(m)
//	kv := reflect.ValueOf(k)
//	vv := reflect.ValueOf(v)
//	if mv.Kind() == reflect.Map &&
//		mv.Type().Key() == kv.Type() &&
//		mv.Type().Elem() == vv.Type() {
//		mv.SetMapIndex(kv, vv)
//	}
//}
//func deleteMap(m, k any) {
//	mv := reflect.ValueOf(m)
//	kv := reflect.ValueOf(k)
//	if mv.Kind() == reflect.Map &&
//		mv.Type().Key() == kv.Type() {
//		mv.SetMapIndex(kv, reflect.Value{})
//	}
//}
//
////	func removeFromMap(m interface{}, key interface{}) {
////	   mapValue := reflect.ValueOf(m)
////	   keyValue := reflect.ValueOf(key)
////	   if (mapValue.Kind() == reflect.Map &&
////	       mapValue.Type().Key() == keyValue.Type()) {
////	           mapValue.SetMapIndex(keyValue, reflect.Value{})
////	   }
////	}
////
////	func main() {
////	   pricesMap := map[string]float64 {
////	       "Kayak": 279, "Lifejacket": 48.95, "Soccer Ball": 19.50,
////	   }
////	   setMap(pricesMap, "Kayak", 100.00)
////	   setMap(pricesMap, "Hat", 10.00)
////	   removeFromMap(pricesMap, "Lifejacket")
////	   for k, v := range pricesMap {
////	       Printfln("Key: %v, Value: %v", k, v)
////	   }
////	}
//func main() {
//	m := map[string]int{
//		"1": 1,
//	}
//	setMap(m, "2", 2)
//	fmt.Println(m)
//	deleteMap(m, "2")
//	fmt.Println(m)
//}
