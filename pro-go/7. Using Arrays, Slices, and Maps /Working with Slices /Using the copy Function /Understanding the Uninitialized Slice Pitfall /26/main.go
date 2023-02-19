package main

import "fmt"

func main() {
	// products := [4]string { "Kayak", "Lifejacket", "Paddle", "Hat"}
	//    allNames := products[1:]
	//    var someNames []string
	//    copy(someNames, allNames)
	//    fmt.Println("someNames:", someNames)
	//    fmt.Println("allNames", allNames)
	lang := [4]string{"java", "go", "py", "rust"}
	gpr := lang[1:]
	var null []string
	copy(null, gpr)
	fmt.Println(null, gpr)
}
