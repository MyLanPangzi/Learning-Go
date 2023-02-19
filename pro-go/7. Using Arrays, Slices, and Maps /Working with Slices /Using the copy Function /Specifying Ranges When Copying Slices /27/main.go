package main

import "fmt"

func main() {
	// products := [4]string { "Kayak", "Lifejacket", "Paddle", "Hat"}
	//    allNames := products[1:]
	//    someNames := []string { "Boots", "Canoe"}
	//    copy(someNames[1:], allNames[2:3])
	//    fmt.Println("someNames:", someNames)
	//    fmt.Println("allNames", allNames)
	lang := [4]string{"java", "go", "py", "rust"}
	gpr := lang[1:]
	ks := []string{"kotlin", "scala"}
	copy(ks[1:], gpr[2:3])
	fmt.Println(ks, gpr)
}
