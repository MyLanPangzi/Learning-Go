package main

import (
	"fmt"
	"math"
)

func main() {
	// when struct implements interface that the receiver is value
	// then pointer and value can assign to interface variable
	var a Abser
	f := Float(-math.Sqrt2)
	v := Vertex8{3, 4}
	a = f
	a = &v
	//a = v //Type does not implement 'Abser' as the 'Abs' method has a pointer receiver

	fmt.Println(a.Abs())
}

//func main() {
//	var a Abser
//	f := MyFloat(-math.Sqrt2)
//	v := Vertex{3, 4}
//
//	a = f  // a MyFloat implements Abser
//	a = &v // a *Vertex implements Abser
//
//	// In the following line, v is a Vertex (not *Vertex)
//	// and does NOT implement Abser.
//	a = v
//
//	fmt.Println(a.Abs())
//}

type Abser interface {
	Abs() float64
}

//	type Abser interface {
//		Abs() float64
//	}

type Float float64

func (f Float) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex8 struct {
	X, Y float64
}

func (v *Vertex8) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
