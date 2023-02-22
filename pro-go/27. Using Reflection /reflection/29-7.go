package main

import (
	"fmt"
	"reflect"
	"strings"
)

//func mapSlice(slice interface{}, mapper interface{}) (mapped []interface{}) {
//    sliceVal := reflect.ValueOf(slice)
//    mapperVal := reflect.ValueOf(mapper)
//    mapped = []interface{} {}
//    if sliceVal.Kind() == reflect.Slice && mapperVal.Kind() == reflect.Func &&
//            mapperVal.Type().NumIn() == 1 &&
//            mapperVal.Type().In(0) == sliceVal.Type().Elem() {
//        for i := 0; i < sliceVal.Len(); i++ {
//            result := mapperVal.Call([]reflect.Value {sliceVal.Index(i)})
//            for _, r := range result {
//                mapped = append(mapped, r.Interface())
//            }
//        }
//    }
//    return
//}
//func main() {
//    names := []string { "Alice", "Bob", "Charlie" }
//    results := mapSlice(names, strings.ToUpper)
//    Printfln("Results: %v", results)
//}

func Map(s, f any) []any {
	sv := reflect.ValueOf(s)
	fv := reflect.ValueOf(f)
	if sv.Kind() != reflect.Slice ||
		fv.Kind() != reflect.Func ||
		fv.Type().NumIn() != 1 ||
		fv.Type().NumOut() != 1 ||
		fv.Type().In(0) != sv.Type().Elem() {
		return nil
	}
	var out = make([]any, sv.Len())
	for i := 0; i < sv.Len(); i++ {
		out[i] = fv.Call([]reflect.Value{sv.Index(i)})[0].Interface()
	}
	return out
}
func Map2[T1 any, T2 any](s []T1, f func(T1) T2) []T2 {
	out := make([]T2, 0, len(s))
	for _, v := range s {
		out = append(out, f(v))
	}
	return out
}
func main() {
	ab := []string{"a", "b"}
	out := Map(ab, strings.ToUpper)
	fmt.Println(out, Map2(ab, strings.ToUpper))
}
