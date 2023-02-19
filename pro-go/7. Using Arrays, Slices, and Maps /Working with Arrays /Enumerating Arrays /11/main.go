package main

import "fmt"

func main() {
	// names := [3]string { "Kayak", "Lifejacket", "Paddle" }
	//    for index, value := range names {
	//        fmt.Println("Index:", index, "Value:", value)
	//    }
	names := [3]string{"hello", "world", "go"}
	for i, name := range names {
		fmt.Println(i, name)
	}
}
