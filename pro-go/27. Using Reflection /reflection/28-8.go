package main

import (
	"fmt"
	"reflect"
)

//	func findAndSplit(slice interface{}, target interface{}) interface{} {
//	   sliceVal := reflect.ValueOf(slice)
//	   targetType := reflect.TypeOf(target)
//	   if (sliceVal.Kind() == reflect.Slice && sliceVal.Type().Elem() == targetType) {
//	       for i := 0; i < sliceVal.Len(); i++ {
//	           if sliceVal.Index(i).Interface() == target {
//	               return sliceVal.Slice(0, i +1)
//	           }
//	       }
//	   }
//	   return slice
//	}
func findAndSplit(slice, target any) any {
	sv := reflect.ValueOf(slice)
	tt := reflect.TypeOf(target)
	if sv.Kind() == reflect.Slice &&
		sv.Type().Elem() == tt &&
		sv.Type().Elem().Comparable() &&
		tt.Comparable() {
		for i := 0; i < sv.Len(); i++ {
			if sv.Index(i).Interface() == target {
				return sv.Slice(0, i+1).Interface()
			}
		}
	}
	return slice
}

func main() {
	s := []string{"1", "2", "3"}
	split := findAndSplit(s, "2")
	fmt.Printf("%v %T\n", split, split)
	ints := []int{1, 2, 3, 4}
	onetowthree := findAndSplit(ints, 3).([]int)
	for i, v := range onetowthree {
		fmt.Println(i, v)
	}
}

//func main() {
//    name := "Alice"
//    city := "London"
//    hobby := "Running"
//    slice := []string { name, city, hobby }
//    //array := [3]string { name, city, hobby}
//    Printfln("Strings: %v", findAndSplit(slice, "London"))
//    numbers := []int {1, 3, 4, 5, 7}
//    Printfln("Numbers: %v", findAndSplit(numbers, 4))
//}
