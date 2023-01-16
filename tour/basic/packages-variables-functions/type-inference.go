package main

import "fmt"

func main() {
	v := 42
	printType(v)
	f := 1.0
	printType(f)
	printType(1i)
}

func printType(v any) {
	fmt.Printf("v is of type %T\n", v)
}

//package main
//
//import "fmt"
//
//func main() {
//	v := 42 // change me!
//	fmt.Printf("v is of type %T\n", v)
//}
