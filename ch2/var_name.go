package main

import "fmt"

func main() {
	_0 := 0_0
	_𝟙 := 20
	π := 3
	ａ := "hello" // Unicode U+FF41
	你好 := "你好"
	fmt.Println(你好)
	fmt.Println(_0)
	fmt.Println(_𝟙)
	fmt.Println(π)
	fmt.Println(ａ)

	ａ = "hello"    // Unicode U+FF41
	a := "goodbye" // standard lowercase a (Unicode U+0061)
	fmt.Println(ａ)
	fmt.Println(a)
}
