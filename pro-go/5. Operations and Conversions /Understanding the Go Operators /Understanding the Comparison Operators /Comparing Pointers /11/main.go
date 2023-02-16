package main

import "fmt"

func main() {

	first := 100
	second := &first
	third := &first
	a := 100
	b := &a
	fmt.Println(*second == *third, *second == *b)
}
