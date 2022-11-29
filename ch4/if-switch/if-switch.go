package main

import (
	"fmt"
	"math/rand"
)

func main() {
	switch n := rand.Intn(10); {
	case n == 0:
		fmt.Println("low")
	case n < 5:
		fmt.Println("big")
	default:
		fmt.Println("good")
	}
}
