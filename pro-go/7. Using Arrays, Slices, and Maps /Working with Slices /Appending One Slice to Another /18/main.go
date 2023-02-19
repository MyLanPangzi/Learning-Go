package main

import (
	"fmt"
	"strings"
)

func main() {
	//  names := make([]string, 3, 6)
	//    names[0] = "Kayak"
	//    names[1] = "Lifejacket"
	//    names[2] = "Paddle"
	//    moreNames := []string { "Hat Gloves"}
	//    appendedNames := append(names, moreNames...)
	//    fmt.Println("appendedNames:", appendedNames)
	names := make([]string, 3, 6)
	names[0] = "hello"
	names = append(names, []string{"world"}...)
	fmt.Println(names, strings.Join(names, ","))
}
