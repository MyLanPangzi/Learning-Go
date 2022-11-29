package main

import "fmt"

func main() {
	x := 10
	if x > 1 {
		fmt.Println(x)
		x := 20
		fmt.Println(x)
	}
	fmt.Println(x)

	a := 10
	if a > 2 {
		a, b := 1, 20
		fmt.Println(a, b)
	}
	fmt.Println(a)

	//fmt := "hello"
	//fmt.Println(fmt)

	fmt.Println(true)
	true := 10
	fmt.Println(true)
}
