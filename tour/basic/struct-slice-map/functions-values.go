package main

import (
	"fmt"
	"math"
)

func main() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(compute(hypot), hypot(5, 12), compute(math.Pow))
}
func compute(f func(float64, float64) float64) float64 {
	return f(3, 4)
}

//package main
//
//import (
//	"fmt"
//	"math"
//)
//
//func compute(fn func(float64, float64) float64) float64 {
//	return fn(3, 4)
//}
//
//func main() {
//	hypot := func(x, y float64) float64 {
//		return math.Sqrt(x*x + y*y)
//	}
//	fmt.Println(hypot(5, 12))
//
//	fmt.Println(compute(hypot))
//	fmt.Println(compute(math.Pow))
//}
