package main

import (
	"fmt"
	"strconv"
)

func main() {
	//product := "Kayak"
	//    for _, character := range product {
	//        fmt.Println("Character:", string(character))
	//    }
	s := "hello 🌞 world"
	for _, v := range s {
		fmt.Println(string(v), strconv.Itoa(int(v)), v)
	}
}
