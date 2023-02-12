package main

import "fmt"

type Integer interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64
}

func Plus[T Integer](i T) T {
	//return i + 1_000 //cannot convert 1_000 (untyped int constant 1000) to T
	return i + 100 //cannot convert 1_000 (untyped int constant 1000) to T

}
func main() {
	r := Plus[uint8](255)
	fmt.Println(r)
}
