package main

//
//import (
//	"fmt"
//	"reflect"
//)
//
//// package main
//// import (
////
////	"reflect"
////	//"strings"
////	//"fmt"
////
//// )
////
////	func invokeFunction(f interface{}, params ...interface{}) {
////	   paramVals := []reflect.Value {}
////	   for _, p := range params {
////	       paramVals = append(paramVals, reflect.ValueOf(p))
////	   }
////	   funcVal := reflect.ValueOf(f)
////	   if (funcVal.Kind() == reflect.Func) {
////	       results := funcVal.Call(paramVals)
////	       for i, r := range results {
////	           Printfln("Result #%v: %v", i, r)
////	       }
////	   }
////	}
////
////	func main() {
////	   names := []string { "Alice", "Bob", "Charlie" }
////	   invokeFunction(Find, names, "London", "Bob")
////	}
//func invoke(f any, args ...any) {
//	fv := reflect.ValueOf(f)
//	if fv.Kind() != reflect.Func {
//		return
//	}
//	in := make([]reflect.Value, len(args))
//	for i := range in {
//		in[i] = reflect.ValueOf(args[i])
//	}
//	out := fv.Call(in)
//	for i, r := range out {
//		fmt.Printf("R#%v %v\n", i, r)
//	}
//}
//func main() {
//	invoke(Find, []string{"1", "2", "3"}, "2")
//}
