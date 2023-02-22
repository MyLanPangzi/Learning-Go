package main

//
//import (
//	"fmt"
//	"reflect"
//)
//
////	func getFieldValues(s interface{}) {
////	   structValue := reflect.ValueOf(s)
////	   if structValue.Kind() == reflect.Struct {
////	       for i := 0; i < structValue.NumField(); i++ {
////	           fieldType := structValue.Type().Field(i)
////	           fieldVal := structValue.Field(i)
////	           Printfln("Name: %v, Type: %v, Value: %v",
////	               fieldType.Name, fieldType.Type, fieldVal)
////	       }
////	   } else {
////	       Printfln("Not a struct")
////	   }
////	}
//func getFieldValues(s any) {
//	sv := reflect.ValueOf(s)
//	if sv.Kind() == reflect.Struct {
//		for i := 0; i < sv.NumField(); i++ {
//			ft := sv.Type().Field(i)
//			fv := sv.Field(i)
//			fmt.Printf("name %v type %v value %v\n", ft.Name, ft, fv)
//		}
//	}
//}
//func main() {
//	p := Product{Name: "apple", Category: "fruit", Price: 10}
//	c := Customer{Name: "hello", City: "hu"}
//	purchase := Purchase{c, p, 100, 10}
//	getFieldValues(purchase)
//}
//
////func main() {
////    product := Product{ Name: "Kayak", Category: "Watersports", Price: 279 }
////    customer := Customer{ Name: "Acme", City: "Chicago" }
////    purchase := Purchase { Customer: customer, Product: product, Total: 279,
////        taxRate: 10 }
////    getFieldValues(purchase)
////}
