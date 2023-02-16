package main

import (
	"fmt"
)

func main() {
	a := 200
	b := 2.1
	c := a + int(b)
	fmt.Println(c, int8(c))
}
