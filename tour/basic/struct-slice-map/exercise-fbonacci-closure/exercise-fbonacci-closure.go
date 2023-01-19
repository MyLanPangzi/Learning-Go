package main

import "fmt"

/*
Let's have some fun with functions.

Implement a fibonacci function that returns a function (a closure) that returns successive fibonacci numbers (0, 1, 1, 2, 3, 5, ...).
*/
func fibonacci() func() int {
	start := 0
	pre := 1
	return func() int {
		// 0+1 0
		// 1 + 0 1
		// 1 + 1 1
		// 2+1  2
		// 3+2  3
		// 5+3 5
		start, pre = start+pre, start
		return pre
	}
}
func main() {
	f := fibonacci()
	for i := 0; i < 11; i++ {
		fmt.Println(f())
	}
}

//package main
//
//import "fmt"
//
//// fibonacci is a function that returns
//// a function that returns an int.
//func fibonacci() func() int {
//}
//
//func main() {
//	f := fibonacci()
//	for i := 0; i < 10; i++ {
//		fmt.Println(f())
//	}
//}
