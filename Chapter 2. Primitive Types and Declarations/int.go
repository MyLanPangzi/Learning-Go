package main

import "fmt"

func main() {
	var i8 int8 = 127
	var i16 int16 = 1<<15 - 1
	var i32 int32 = 1<<31 - 1
	var i64 int64 = 1<<63 - 1
	var u8 uint8 = 1<<8 - 1
	var u16 uint16 = 1<<16 - 1
	var u32 uint32 = 1<<32 - 1
	var u64 uint64 = 1<<64 - 1
	fmt.Println(i8, i16, i32, i64, u8, u16, u32, u64)

	var i int = 1<<63 - 1
	var ui uint = 1<<64 - 1
	fmt.Println(i, ui)

	var b byte = 1<<8 - 1
	fmt.Println(b)

	var x int = 10
	x *= 2
	fmt.Println(x)
}
