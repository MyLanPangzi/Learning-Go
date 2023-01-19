package main

import "fmt"

func main() {
	a := make([]int, 5)
	printSlice2("a", a)

	b := make([]int, 0, 5)
	printSlice2("b", b)

	c := b[:2]
	printSlice2("c", c)

	d := c[2:5]
	printSlice2("d", d)
}

func printSlice2(s string, a []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(a), cap(a), a)
}

//package main
//
//import "fmt"
//
//func main() {
//	a := make([]int, 5)
//	printSlice("a", a)
//
//	b := make([]int, 0, 5)
//	printSlice("b", b)
//
//	c := b[:2]
//	printSlice("c", c)
//
//	d := c[2:5]
//	printSlice("d", d)
//}
//
//func printSlice(s string, x []int) {
//	fmt.Printf("%s len=%d cap=%d %v\n",
//		s, len(x), cap(x), x)
//}
