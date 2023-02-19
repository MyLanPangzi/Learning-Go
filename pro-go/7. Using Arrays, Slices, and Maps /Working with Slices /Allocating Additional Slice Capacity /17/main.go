package main

import "fmt"

func main() {
	// names := make([]string, 3, 6)
	//    names[0] = "Kayak"
	//    names[1] = "Lifejacket"
	//    names[2] = "Paddle"
	//    appendedNames := append(names, "Hat", "Gloves")
	//    names[0] = "Canoe"
	//    fmt.Println("names:",names)
	//    fmt.Println("appendedNames:", appendedNames)
	names := make([]string, 3, 6)
	names[0] = "hello"
	names[1] = "world"
	names[2] = "go"
	appendNames := append(names, "java", "kotlin")
	names[0] = "scala"
	fmt.Println(names, appendNames)
}
