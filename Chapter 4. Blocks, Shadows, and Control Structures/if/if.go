package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x := rand.Intn(10)
	fmt.Println(x)
	if x == 0 {
		fmt.Println("low")
	} else if x > 5 {
		fmt.Println("big")
	} else {
		fmt.Println("good")
	}

	if n := rand.Intn(10); n == 0 {
		fmt.Println("low")
	}
	//fmt.Println(n)
}
