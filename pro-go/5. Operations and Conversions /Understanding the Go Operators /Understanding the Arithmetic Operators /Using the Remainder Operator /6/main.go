package main

import (
	"fmt"
	"math"
)

func main() {
	negResult := -3 % 2
	fmt.Println(3%2, negResult, math.Abs(float64(negResult)))
}
