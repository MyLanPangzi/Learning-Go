package main

//
//import (
//	"fmt"
//	"reflect"
//	"strings"
//)
//
////	func makeMapperFunc(mapper interface{}) interface{} {
////	   mapVal := reflect.ValueOf(mapper)
////	   if mapVal.Kind() == reflect.Func && mapVal.Type().NumIn() == 1 &&
////	           mapVal.Type().NumOut() == 1  {
////	       inType := reflect.SliceOf( mapVal.Type().In(0))
////	       inTypeSlice := []reflect.Type { inType }
////	       outType := reflect.SliceOf( mapVal.Type().Out(0))
////	       outTypeSlice := []reflect.Type { outType }
////	       funcType := reflect.FuncOf(inTypeSlice, outTypeSlice, false)
////	       funcVal := reflect.MakeFunc(funcType,
////	               func (params []reflect.Value) (results []reflect.Value) {
////	           srcSliceVal := params[0]
////	           resultsSliceVal := reflect.MakeSlice(outType, srcSliceVal.Len(), 10)
////	           for i := 0; i < srcSliceVal.Len(); i++ {
////	               r := mapVal.Call([]reflect.Value { srcSliceVal.Index(i)})
////	               resultsSliceVal.Index(i).Set(r[0])
////	           }
////	           results = []reflect.Value { resultsSliceVal }
////	           return
////	       })
////	       return funcVal.Interface()
////	   }
////	   Printfln("Unexpected types")
////	   return nil
////	}
//func makeMapperFunc(f any) any {
//	fv := reflect.ValueOf(f)
//	if !(fv.Kind() == reflect.Func &&
//		fv.Type().NumIn() == 1) || fv.Type().NumOut() != 1 {
//		return nil
//	}
//	outType := reflect.SliceOf(fv.Type().Out(0))
//	ft := reflect.FuncOf(
//		[]reflect.Type{reflect.SliceOf(fv.Type().In(0))},
//		[]reflect.Type{outType},
//		false,
//	)
//	return reflect.MakeFunc(ft, func(args []reflect.Value) []reflect.Value {
//		out := reflect.MakeSlice(outType, 0, args[0].Len())
//		for i := 0; i < args[0].Len(); i++ {
//			out = reflect.Append(out, fv.Call([]reflect.Value{args[0].Index(i)})...)
//		}
//		return []reflect.Value{out}
//	}).Interface()
//}
//func main() {
//	lower := makeMapperFunc(strings.ToLower).(func([]string) []string)
//	fmt.Println(lower([]string{"A", "B"}))
//	double := makeMapperFunc(func(i int) int { return i * 2 }).(func([]int) []int)
//	fmt.Println(double([]int{1, 2}))
//	addOne := makeMapperFunc2(func(i int) int { return i + 1 })
//	fmt.Println(addOne([]int{1, 2}))
//	upper := makeMapperFunc2(strings.ToUpper)
//	fmt.Println(upper([]string{"a", "b"}))
//
//}
//func makeMapperFunc2[T1 any, T2 any](f func(T1) T2) func([]T1) []T2 {
//	return func(in []T1) []T2 {
//		out := make([]T2, 0, len(in))
//		for _, v := range in {
//			out = append(out, f(v))
//		}
//		return out
//	}
//}
//
////func main() {
////    lowerStringMapper := makeMapperFunc(strings.ToLower).(func([]string)[]string)
////    names := []string { "Alice", "Bob", "Charlie" }
////    results := lowerStringMapper(names)
////    Printfln("Lowercase Results: %v", results)
////    incrementFloatMapper := makeMapperFunc(func (val float64) float64 {
////        return val + 1
////    }).(func([]float64)[]float64)
////    prices := []float64 { 279, 48.95, 19.50}
////    floatResults := incrementFloatMapper(prices)
////    Printfln("Increment Results: %v", floatResults)
////    floatToStringMapper := makeMapperFunc(func (val float64) string {
////        return fmt.Sprintf("$%.2f", val)
////    }).(func([]float64)[]string)
////    Printfln("Price Results: %v", floatToStringMapper(prices))
////}
