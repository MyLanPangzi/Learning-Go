package main

import "fmt"

func main() {
	first := 100
	second := &first
	first++
	fmt.Println(first, second)
}
