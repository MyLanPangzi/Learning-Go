package main

//
//import "reflect"
//
//func printDetails(values ...any) {
//	for _, e := range values {
//		v := reflect.ValueOf(e)
//		switch v.Kind() {
//		case reflect.Bool:
//			Printfln("Bool %v", v.Bool())
//		case reflect.Int:
//			Printfln("Int %v", v.Int())
//		case reflect.Float32, reflect.Float64:
//			Printfln("Float %v", v.Float())
//		case reflect.String:
//			Printfln("String %v", v.String())
//		case reflect.Ptr:
//			elem := v.Elem()
//			if elem.Kind() == reflect.Int {
//				Printfln("Int Ptr %v", elem.Int())
//			}
//		default:
//			Printfln("Other %v", v.String())
//		}
//	}
//}
//func main() {
//	p := Product{}
//	i := 100
//	printDetails(true, 10, 2.0, "Alice", &i, p)
//}
