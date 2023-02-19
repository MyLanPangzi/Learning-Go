package main

import "fmt"

func main() {
	// products := [4]string { "Kayak", "Lifejacket", "Paddle", "Hat"}
	//    someNames := products[1:3]
	//    allNames := products[:]
	//    fmt.Println("someNames:", someNames)
	//    fmt.Println("someNames len:", len(someNames), "cap:", cap(someNames))
	//    fmt.Println("allNames", allNames)
	//    fmt.Println("allNames len", len(allNames), "cap:", cap(allNames))
	lang := [...]string{"java", "go", "python", "rust"}
	goPy := lang[1:3]
	goPy = append(goPy, "kotlin")
	goPy = append(goPy, "scala")
	all := lang[:]
	fmt.Println(goPy, len(goPy), cap(goPy))
	fmt.Println(all, len(all), cap(all))
}
