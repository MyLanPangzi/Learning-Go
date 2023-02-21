package _7

//
//import "reflect"
//
//func selectValue(data any, index int) any {
//	v := reflect.ValueOf(data)
//	if v.Kind() == reflect.Slice {
//		return v.Index(index).Interface()
//	}
//	return nil
//}
//func selectValue2[T any](data []T, index int) T {
//	return data[index]
//}
//func main() {
//	names := []string{"Alice", "Bob", "Charlie"}
//	v := selectValue(names, 1).(string)
//	Printfln("selected : %v", v)
//	Printfln("selected : %v", selectValue2(names, 1))
//}
