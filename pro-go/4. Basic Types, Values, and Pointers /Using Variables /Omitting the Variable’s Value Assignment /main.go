package main

import "fmt"

func main() {
	var p float32
	var p2 float64
	fmt.Printf("%v %v, %T %T\n", p, p2, p, p2)
	p = 10
	p2 = 10
	fmt.Printf("%v %v, %T %T\n", p, p2, p, p2)
}
