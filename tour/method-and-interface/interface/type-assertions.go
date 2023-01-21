package main

import "fmt"

func main() {
	var i any = "hello"
	s := i.(string)
	fmt.Println(s)

	f, ok := i.(float64)
	fmt.Println(f, ok)
	f = i.(float64)
	fmt.Println(f)
}

//package main
//
//import "fmt"
//
//func main() {
//	var i interface{} = "hello"
//
//	s := i.(string)
//	fmt.Println(s)
//
//	s, ok := i.(string)
//	fmt.Println(s, ok)
//
//	f, ok := i.(float64)
//	fmt.Println(f, ok)
//
//	f = i.(float64) // panic
//	fmt.Println(f)
//}
