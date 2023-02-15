package main

import "fmt"

func main() {
	first := 100
	second := &first
	third := second
	first++
	*second++
	*third++
	fmt.Println(first, *second, *third)
}
