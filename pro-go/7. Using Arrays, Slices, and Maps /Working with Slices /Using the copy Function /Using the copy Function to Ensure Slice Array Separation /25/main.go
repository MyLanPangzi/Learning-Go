package main

import "fmt"

func main() {
	//  products := [4]string { "Kayak", "Lifejacket", "Paddle", "Hat"}
	//    allNames := products[1:]
	//    someNames := make([]string, 2)
	//    copy(someNames, allNames)
	//    fmt.Println("someNames:", someNames)
	//    fmt.Println("allNames", allNames)
	lang := [...]string{"java", "go", "py", "rust"}
	gpr := lang[1:]
	gp := make([]string, 2)
	copy(gp, gpr)
	fmt.Println(gp, gpr)
}
