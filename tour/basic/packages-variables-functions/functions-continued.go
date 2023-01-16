package main

import "fmt"

// package main
//
// import "fmt"
//
//	func add(x, y int) int {
//		return x + y
//	}
//
//	func main() {
//		fmt.Println(add(42, 13))
//	}
func add2(x, y int) int {
	return x + y
}
func main() {
	fmt.Println(add2(2, 3))
}
