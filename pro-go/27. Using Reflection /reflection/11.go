package main

//
//import (
//	"reflect"
//)
//
//var intPtrType = reflect.TypeOf((*int)(nil))
//
//func printDetails(values ...any) {
//	for _, e := range values {
//		t := reflect.TypeOf(e)
//		v := reflect.ValueOf(e)
//		if t.Kind() == reflect.Pointer && t.Elem().Kind() == reflect.Int {
//			Printfln("Pointer to Int:%v", v.Elem().Int())
//			continue
//		}
//		//if t == intPtrType {
//		//	Printfln("Pointer to Int:%v\n", v.Elem().Int())
//		//	continue
//		//}
//		switch v.Kind() {
//		case reflect.Bool:
//			var val bool = v.Bool()
//			Printfln("Bool: %v", val)
//		case reflect.Int:
//			var val int64 = v.Int()
//			Printfln("Int: %v", val)
//		case reflect.Float32, reflect.Float64:
//			var val float64 = v.Float()
//			Printfln("Float: %v", val)
//		case reflect.String:
//			var val string = v.String()
//			Printfln("String: %v", val)
//		default:
//			Printfln("Other: %v", v.String())
//		}
//	}
//}
//
//func main() {
//	product := Product{
//		Name: "Kayak", Category: "Watersports", Price: 279,
//	}
//	number := 100
//	printDetails(true, 10, 23.30, "Alice", &number, product)
//}
