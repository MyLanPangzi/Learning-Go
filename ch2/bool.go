package main

import "fmt"

func main() {
	b := false
	fmt.Println(b)
	b = true
	a, b := true, true
	fmt.Println(b, a)

	var flag bool // no value assigned, set to false
	var isAwesome = true

	fmt.Println(flag, isAwesome)
}
