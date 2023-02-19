package main

import (
	"fmt"
	"reflect"
)

func main() {
	// p1 := []string { "Kayak", "Lifejacket", "Paddle", "Hat"}
	//    p2 := p1
	//    fmt.Println("Equal:", p1 == p2)

	s1 := []string{"1", "2"}
	s2 := []string{"1", "2"}
	fmt.Println(reflect.DeepEqual(s1, s2))
}
