package main

import "fmt"

type BuiltInOrdered interface {
	~string | ~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~uintptr | ~float64 | ~float32
}

func Min[T BuiltInOrdered](x, y T) T {
	if x < y {
		return x
	}
	return y
}

func main() {
	x := 10
	y := 20
	fmt.Println(Min(x, y))
	type MyInt int
	var a MyInt = 10
	var b MyInt = 20
	fmt.Println(Min(a, b))
}
