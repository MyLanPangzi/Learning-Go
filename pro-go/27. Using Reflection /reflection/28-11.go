package main

//
//import (
//	"fmt"
//	"reflect"
//)
//
////func printMapContents(m interface{}) {
////    mapValue := reflect.ValueOf(m)
////    if (mapValue.Kind() == reflect.Map) {
////        for _, keyVal := range mapValue.MapKeys() {
////            reflectedVal := mapValue.MapIndex(keyVal)
////            Printfln("Map Key: %v, Value: %v", keyVal, reflectedVal)
////        }
////    } else {
////        Printfln("Not a map")
////    }
////}
////func main() {
////    pricesMap := map[string]float64 {
////        "Kayak": 279, "Lifejacket": 48.95, "Soccer Ball": 19.50,
////    }
////    printMapContents(pricesMap)
////}
//
//func main() {
//	m := map[string]int{
//		"1": 1,
//		"2": 2,
//	}
//	mv := reflect.ValueOf(m)
//	if mv.Kind() == reflect.Map {
//		for _, k := range mv.MapKeys() {
//			fmt.Println(k, mv.MapIndex(k))
//		}
//
//		iter := mv.MapRange()
//		for iter.Next() {
//			fmt.Println(iter.Key(), iter.Value())
//		}
//	}
//
//}
