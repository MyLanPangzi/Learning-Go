package main

import "fmt"

func main() {
	first := 100
	second := &first
	*second++
	fmt.Println(first, *second)
}
