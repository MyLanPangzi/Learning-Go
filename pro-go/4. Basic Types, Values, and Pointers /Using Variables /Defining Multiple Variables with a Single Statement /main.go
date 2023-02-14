package main

import "fmt"

func main() {
	var p, t = 2.0, 0.2
	fmt.Println(p + t)
	fmt.Printf("%v %v %T %T\n", p, t, p, t)
	var p2, t2 float64
	p2 = 2
	t2 = 0.2
	fmt.Printf("%v %v %T %T\n", p2, t2, p2, t2)

}
