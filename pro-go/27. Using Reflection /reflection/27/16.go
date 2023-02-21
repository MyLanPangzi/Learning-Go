package _7

//
//import "reflect"
//
//func setAll(src any, targets ...any) {
//	v := reflect.ValueOf(src)
//	for _, t := range targets {
//		tv := reflect.ValueOf(t)
//		if tv.Kind() == reflect.Pointer && tv.Elem().Type() == v.Type() && tv.Elem().CanSet() {
//			tv.Elem().Set(v)
//		}
//	}
//}
//func main() {
//	n := "hello"
//	p := 100
//	c := "london"
//	setAll("New", &n, &p, &c)
//	setAll(10, &n, &p, &c)
//	for _, v := range []any{n, p, c} {
//		Printfln("v %v", v)
//	}
//}
