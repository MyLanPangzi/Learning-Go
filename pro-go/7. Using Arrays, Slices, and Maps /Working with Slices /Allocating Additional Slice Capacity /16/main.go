package main

import "fmt"

func main() {
	// names := make([]string, 3, 6)
	//    names[0] = "Kayak"
	//    names[1] = "Lifejacket"
	//    names[2] = "Paddle"
	//    fmt.Println("len:", len(names))
	//    fmt.Println("cap:", cap(names))
	names := make([]string, 3, 6)
	names[0] = "hello"
	names[1] = "world"
	names[2] = "go"
	fmt.Println(len(names), cap(names))
}
