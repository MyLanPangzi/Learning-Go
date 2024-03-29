package main

//
//import (
//	"reflect"
//)
//
////	func setFieldValue(s interface{}, newVals map[string]interface{}) {
////	   structValue := reflect.ValueOf(s)
////	   if (structValue.Kind() == reflect.Ptr &&
////	           structValue.Elem().Kind() == reflect.Struct) {
////	       for name, newValue := range newVals {
////	           fieldVal := structValue.Elem().FieldByName(name)
////	           if (fieldVal.CanSet()) {
////	               fieldVal.Set(reflect.ValueOf(newValue))
////	           } else if (fieldVal.CanAddr()) {
////	               ptr := fieldVal.Addr()
////	               if (ptr.CanSet()) {
////	                   ptr.Set(reflect.ValueOf(newValue))
////	               } else {
////	                   Printfln("Cannot set field via pointer")
////	               }
////	           } else {
////	               Printfln("Cannot set field")
////	           }
////	       }
////	   } else {
////	       Printfln("Not a pointer to a struct")
////	   }
////	}
//func setFieldValue(s any, vals map[string]any) {
//	sv := reflect.ValueOf(s)
//	if sv.Kind() != reflect.Pointer || sv.Elem().Kind() != reflect.Struct {
//		return
//	}
//	for k, v := range vals {
//		fv := sv.Elem().FieldByName(k)
//		if fv.CanSet() {
//			fv.Set(reflect.ValueOf(v))
//			continue
//		}
//		//if fv.CanAddr() {
//		//	fmt.Printf("%v %v\n", k, fv.Type())
//		//	ptr := fv.Addr()
//		//	if ptr.CanSet() {
//		//		ptr.Set(reflect.ValueOf(v))
//		//	}
//		//}
//	}
//}
//
//func getFieldValues(s interface{}) {
//	structValue := reflect.ValueOf(s)
//	if structValue.Kind() == reflect.Struct {
//		for i := 0; i < structValue.NumField(); i++ {
//			fieldType := structValue.Type().Field(i)
//			fieldVal := structValue.Field(i)
//			Printfln("Name: %v, Type: %v, Value: %v",
//				fieldType.Name, fieldType.Type, fieldVal)
//		}
//	} else {
//		Printfln("Not a struct")
//	}
//}
//
//type Person struct {
//	Pet *string
//}
//
//func main() {
//	p := Product{"apple", "fruit", 1}
//	c := Customer{"hello", "hu"}
//	purchase := Purchase{c, p, 100, 10}
//	setFieldValue(&purchase, map[string]any{
//		"City":     "London",
//		"Category": "food",
//		"Total":    200.0,
//	})
//	getFieldValues(purchase)
//	person := Person{}
//	s := "pupy"
//	setFieldValue(&person, map[string]any{
//		"Pet": &s,
//	})
//	getFieldValues(person)
//
//}
//
////func main() {
////    product := Product{ Name: "Kayak", Category: "Watersports", Price: 279 }
////    customer := Customer{ Name: "Acme", City: "Chicago" }
////    purchase := Purchase { Customer: customer, Product: product, Total: 279,
////        taxRate: 10 }
////    setFieldValue(&purchase, map[string]interface{} {
////        "City": "London", "Category": "Boats", "Total": 100.50,
////    })
////    getFieldValues(purchase)
////}
