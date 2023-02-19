package main

import "fmt"

func main() {
	//   names := make([]string, 3)
	//    names[0] = "Kayak"
	//    names[1] = "Lifejacket"
	//    names[2] = "Paddle"
	//    fmt.Println(names)
	names := make([]string, 3)
	names[0] = "hello"
	names[1] = "world"
	names[2] = "go"
	names2 := []string{"java", "kotlin", "scala"}
	fmt.Println(names, names2)
}
