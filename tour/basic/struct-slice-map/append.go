package main

import "fmt"

func main() {
	var s []int
	printSlice3(s)

	s = append(s, 0)
	printSlice3(s)

	s = append(s, 1)
	printSlice3(s)

	//s = append(s, 2, 3, 4)
	s = append(s, 2)
	printSlice3(s)

	s = append(s, 3, 4)
	printSlice3(s)
}

func printSlice3(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

//package main
//
//import "fmt"
//
//func main() {
//	var s []int
//	printSlice(s)
//
//	// append works on nil slices.
//	s = append(s, 0)
//	printSlice(s)
//
//	// The slice grows as needed.
//	s = append(s, 1)
//	printSlice(s)
//
//	// We can add more than one element at a time.
//	s = append(s, 2, 3, 4)
//	printSlice(s)
//}
//
//func printSlice(s []int) {
//	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
//}
