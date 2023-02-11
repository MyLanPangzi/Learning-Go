package main

import (
	"fmt"
	"math"
)

type Pair[T fmt.Stringer] struct {
	Val1 T
	Val2 T
}
type Differ[T any] interface {
	fmt.Stringer
	Diff(T) float64
}
type Point2D struct {
	X, Y int
}

func (p Point2D) String() string {
	return fmt.Sprintf("{%d,%d}", p.X, p.Y)
}

func (p Point2D) Diff(p2 Point2D) float64 {
	x := p.X - p2.X
	y := p.Y - p2.Y
	return math.Sqrt(float64(x*x) + float64(y*y))
}

type Point3D struct {
	X, Y, Z int
}

func (p Point3D) String() string {
	return fmt.Sprintf("{%d,%d,%d}", p.X, p.Y, p.Z)
}

func (p Point3D) Diff(p2 Point3D) float64 {
	x := p.X - p2.X
	y := p.Y - p2.Y
	z := p.Z - p2.Z
	return math.Sqrt(float64(x*x) + float64(y*y) + float64(z*z))
}

func FindCloser[T Differ[T]](pair1, pair2 Pair[T]) Pair[T] {
	d1 := pair1.Val1.Diff(pair1.Val2)
	d2 := pair1.Val1.Diff(pair1.Val2)
	if d1 < d2 {
		return pair1
	}
	return pair2
}
func main() {
	pair2Da := Pair[Point2D]{Val1: Point2D{1, 1}, Val2: Point2D{5, 5}}
	pair2Db := Pair[Point2D]{Val1: Point2D{10, 10}, Val2: Point2D{15, 5}}
	closer := FindCloser(pair2Da, pair2Db)
	fmt.Println(closer)
	p3da := Pair[Point3D]{Val1: Point3D{1, 1, 10}, Val2: Point3D{5, 5, 0}}
	p3db := Pair[Point3D]{Val1: Point3D{10, 10, 10}, Val2: Point3D{11, 5, 10}}
	c2 := FindCloser(p3da, p3db)
	fmt.Println(c2)

	//FindCloser(pair2Da, p3da) // type Pair[Point3D] of p3da does not match inferred type Pair[Point2D] for Pair[T]
	//FindCloser(Pair[StringerString]{"a", "b"}, Pair[StringerString]{"a", "b"}) //StringerString does not implement Differ[StringerString] (missing Diff method)
}

type StringerString string

func (s StringerString) String() string {
	return string(s)
}
