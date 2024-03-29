package main

import "fmt"

func main() {
	ints := []int{1, 2, 3, -1, -2}
	fmt.Println(Index(ints, 10))
	fmt.Println(Index(ints, 1))

	strings := []string{"hello", "world"}
	fmt.Println(Index(strings, "go"))
	fmt.Println(Index(strings, "world"))
}

func Index[T comparable](s []T, x T) int {
	for i, v := range s {
		if v == x {
			return i
		}
	}
	return -1
}

//func main() {
//	// Index works on a slice of ints
//	si := []int{10, 20, 15, -10}
//	fmt.Println(Index(si, 15))
//
//	// Index also works on a slice of strings
//	ss := []string{"foo", "bar", "baz"}
//	fmt.Println(Index(ss, "hello"))
//}
