package main

import "fmt"

func main() {
	//names := []string {"Kayak", "Lifejacket", "Paddle"}
	//    appendedNames := append(names, "Hat", "Gloves")
	//    names[0] = "Canoe"
	//    fmt.Println("names:", names)
	//    fmt.Println("appendedNames:", appendedNames)
	names := []string{"hello", "world", "go"}
	appendNames := append(names, "java", "kotlin")
	fmt.Println(names, appendNames)
}
