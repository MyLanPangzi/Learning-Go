package main

import (
	"fmt"
	"reflect"
)

func main() {
	//  names := [3]string { "Kayak", "Lifejacket", "Paddle" }
	//    moreNames := [3]string { "Kayak", "Lifejacket", "Paddle" }
	//    same := names == moreNames
	//    fmt.Println("comparison:", same)
	names1 := [3]string{"hello", "world", "go"}
	names2 := [3]string{"hello", "world", "go"}
	fmt.Println(names2 == names1, reflect.DeepEqual(names1, names2))

}
