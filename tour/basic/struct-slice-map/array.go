package main

import "fmt"

func main() {
	var a [2]string
	fmt.Printf("%q\n", a)
	a[0] = "hello"
	a[1] = "world"
	fmt.Println(a, a[0], a[1])
}

//package main
//
//import "fmt"
//
//func main() {
//	var a [2]string
//	a[0] = "Hello"
//	a[1] = "World"
//	fmt.Println(a[0], a[1])
//	fmt.Println(a)
//
//	primes := [6]int{2, 3, 5, 7, 11, 13}
//	fmt.Println(primes)
//}
