package main

import "fmt"

func main() {
	p, t, i, _ := 3.0, 0.3, true, true
	var _ = "hello"
	fmt.Println(p, t, i, int(p), []byte("hello"), []uint8("hello"), []int32("hello"), []rune("hello"))

}
