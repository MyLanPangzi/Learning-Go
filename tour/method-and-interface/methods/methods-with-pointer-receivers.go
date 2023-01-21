package main

import (
	"fmt"
	"math"
)

func main() {
	v := &Vertex7{3, 4}
	fmt.Printf("Before scaling: %+v, Abs:%v \n", v, v.Abs())
	v.Scale(5)
	fmt.Printf("After scaling: %v, Abs:%v \n", v, v.Abs())
}

//	func main() {
//		v := &Vertex{3, 4}
//		fmt.Printf("Before scaling: %+v, Abs: %v\n", v, v.Abs())
//		v.Scale(5)
//		fmt.Printf("After scaling: %+v, Abs: %v\n", v, v.Abs())
//	}
type Vertex7 struct {
	X, Y float64
}

func (v *Vertex7) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
func (v *Vertex7) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

//type Vertex struct {
//	X, Y float64
//}
//
//func (v *Vertex) Scale(f float64) {
//	v.X = v.X * f
//	v.Y = v.Y * f
//}
//
//func (v *Vertex) Abs() float64 {
//	return math.Sqrt(v.X*v.X + v.Y*v.Y)
//}
//
