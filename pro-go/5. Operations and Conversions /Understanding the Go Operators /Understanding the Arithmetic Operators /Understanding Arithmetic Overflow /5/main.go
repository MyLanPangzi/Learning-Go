package main

import (
	"fmt"
	"math"
)

func main() {
	maxInt64 := math.MaxInt64
	maxFloat64 := math.MaxFloat64
	fmt.Println(maxInt64*2, maxFloat64*2, math.IsInf(maxFloat64*2, 0))
}
