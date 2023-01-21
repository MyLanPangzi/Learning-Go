package main

import "fmt"

func main() {
	var i any
	desc5(i)

	i = 42
	desc5(i)

	i = "hello"
	desc5(i)
}
func desc5(i any) {
	fmt.Printf("%v %T\n", i, i)
}

//func main() {
//	var i interface{}
//	describe(i)
//
//	i = 42
//	describe(i)
//
//	i = "hello"
//	describe(i)
//}
//
//func describe(i interface{}) {
//	fmt.Printf("(%v, %T)\n", i, i)
//}
