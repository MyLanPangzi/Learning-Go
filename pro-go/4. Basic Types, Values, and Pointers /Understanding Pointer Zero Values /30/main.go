package main

import "fmt"

func main() {
	first := 100
	var second *int
	//fmt.Println(*second) //panic: runtime error: invalid memory address or nil pointer dereference
	second = &first
	fmt.Println(second == nil)
}
