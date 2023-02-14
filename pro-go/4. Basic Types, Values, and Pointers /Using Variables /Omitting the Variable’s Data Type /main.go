package main

import "fmt"

func main() {
	var p = 2.0
	var p2 = p
	fmt.Println(p, p2)
	fmt.Printf("%T %T %T\n", p, p2, 'A')
}
