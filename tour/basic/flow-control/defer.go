package main

import "fmt"

func main() {
	defer fmt.Println("world")
	fmt.Println("hello")
	i := 10
	defer fmt.Println(i)
}

//package main
//
//import "fmt"
//
//func main() {
//	defer fmt.Println("world")
//
//	fmt.Println("hello")
//}
