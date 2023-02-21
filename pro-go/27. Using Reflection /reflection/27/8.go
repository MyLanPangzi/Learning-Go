package _7

//
//import (
//	"fmt"
//	"reflect"
//	"strings"
//)
//
//func printDetails(values ...any) {
//	for _, elem := range values {
//		fieldDetails := []string{}
//		elemType := reflect.TypeOf(elem)
//		elemValue := reflect.ValueOf(elem)
//		if elemType.Kind() != reflect.Struct {
//			Printfln("%v: %v", elemType.Name(), elemValue)
//			continue
//		}
//		for i := 0; i < elemType.NumField(); i++ {
//			name := elemType.Field(i).Name
//			val := elemValue.Field(i)
//			fieldDetails = append(fieldDetails, fmt.Sprintf("%v: %v", name, val))
//		}
//		Printfln("%v:%v", elemType.Name(), strings.Join(fieldDetails, ", "))
//	}
//}
//
//type Payment struct {
//	Currency string
//	Amount   float64
//}
//
//func main() {
//	apple := Product{
//		Name:     "apple",
//		Category: "fruit",
//		Price:    1,
//	}
//	c := Customer{
//		Name: "hello",
//		City: "hu",
//	}
//	p := Payment{
//		Currency: "CNY",
//		Amount:   1,
//	}
//	printDetails(apple, c, p, 10, true)
//}
