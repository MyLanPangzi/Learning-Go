package main

import "fmt"

func main() {
	//products := [4]string { "Kayak", "Lifejacket", "Paddle", "Hat"}
	//    someNames := products[1:3]
	//    allNames := products[:]
	//    fmt.Println("someNames:", someNames)
	//    fmt.Println("allNames", allNames)

	lang := [...]string{"java", "go", "python", "rust"}
	goPython := lang[1:3]
	all := lang[:]
	fmt.Println(goPython, all)
}
