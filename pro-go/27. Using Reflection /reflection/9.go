package main

//
//import "reflect"
//
//func getTypePath(t reflect.Type) string {
//	if t.PkgPath() == "" {
//		return "(built-in)"
//	}
//	return t.PkgPath()
//}
//func printDetails(vals ...any) {
//	for _, e := range vals {
//		t := reflect.TypeOf(e)
//		Printfln("Name:%v,PkgPath:%v,kind:%v", t.Name(), getTypePath(t), t.Kind())
//	}
//}
//
//type Payment struct {
//	Currency string
//	Amount   float64
//}
//
//func main() {
//	a := Product{
//		Name:     "apple",
//		Category: "fruit",
//		Price:    0,
//	}
//	c := Customer{
//		Name: "hello",
//		City: "w",
//	}
//	p := Payment{
//		Currency: "$",
//		Amount:   0,
//	}
//	printDetails(a, c, p, 10, true)
//}
