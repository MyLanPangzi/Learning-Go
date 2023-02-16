package main

import (
	"fmt"
	"math"
)

func main() {
	a := 2
	b := 3.5
	c := a + int(math.Round(b))
	d := a + int(math.Ceil(b))
	e := a + int(math.Floor(b))
	f := a + int(math.RoundToEven(3.4))
	fmt.Println(c, d, e, f, math.RoundToEven(3.4))
}
