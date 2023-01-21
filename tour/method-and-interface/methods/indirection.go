package main

import "fmt"

func main() {
	v := Vertex5{3, 4}
	v.Scale(2)
	ScaleFunc(&v, 10)

	p := &Vertex5{3, 4}
	p.Scale(3)
	ScaleFunc(p, 8)

	fmt.Println(v, p)
}

//func main() {
//	v := Vertex{3, 4}
//	v.Scale(2)
//	ScaleFunc(&v, 10)
//
//	p := &Vertex{4, 3}
//	p.Scale(3)
//	ScaleFunc(p, 8)
//
//	fmt.Println(v, p)
//}

type Vertex5 struct {
	X, Y float64
}

func (v *Vertex5) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}
func ScaleFunc(v *Vertex5, f float64) {
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
//func ScaleFunc(v *Vertex, f float64) {
//	v.X = v.X * f
//	v.Y = v.Y * f
//}
//
