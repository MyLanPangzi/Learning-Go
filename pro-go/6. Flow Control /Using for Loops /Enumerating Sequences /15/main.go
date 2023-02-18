package main

import "fmt"

func main() {
	//  product := "Kayak"
	//    for index, character := range product {
	//        fmt.Println("Index:", index, "Character:", string(character))
	//    }
	s := "hello 🌞 world"
	for i, c := range s {
		fmt.Println("i", i, "c", string(c))
	}
}
