package main

import "fmt"

func main() {
	//  names := [3]string { "Kayak", "Lifejacket", "Paddle" }
	//    otherArray := names
	//    names[0] = "Canoe"
	//    fmt.Println("names:", names)
	//    fmt.Println("otherArray:", otherArray)
	strings := [...]string{"hello", "world"}
	s := strings
	strings[0] = "java"
	sp := &strings
	fmt.Println(strings, s, *sp)
}
