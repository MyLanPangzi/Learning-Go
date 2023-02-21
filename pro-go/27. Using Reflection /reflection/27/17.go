package _7

//
//import (
//	"fmt"
//	"reflect"
//)
//
//func contains(slice any, target any) bool {
//	v := reflect.ValueOf(slice)
//	//tt := reflect.TypeOf(target)
//	if v.Kind() == reflect.Slice {
//		for i := 0; i < v.Len(); i++ {
//			if reflect.DeepEqual(v.Index(i).Interface(), target) {
//				return true
//			}
//			//value := v.Index(i)
//			//if value.Type().Comparable() && tt.Comparable() && value.Interface() == target {
//			//	return true
//			//}
//		}
//	}
//	return false
//}
//func contains2[T comparable](slice []T, target T) bool {
//	for _, v := range slice {
//		if v == target {
//			return true
//		}
//	}
//	return false
//}
//
//func main() {
//	s := []string{"1", "2", "3"}
//	fmt.Println(contains(s, "1"), contains(s, "4"))
//	ss := [][]string{
//		{"1", "2", "3"},
//	}
//	fmt.Println(contains(ss, s))
//}
