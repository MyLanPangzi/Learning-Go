package main

import "fmt"

func main() {
	first := 100
	secod := &first
	third := &first
	a := 100
	b := &a
	fmt.Println(secod == third, secod == b)
}
