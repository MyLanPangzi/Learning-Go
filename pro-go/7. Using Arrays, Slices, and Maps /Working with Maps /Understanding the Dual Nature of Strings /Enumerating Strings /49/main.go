package main

import "fmt"

func main() {
	//var price = "€48.95"
	//    for index, char := range price {
	//        fmt.Println(index, char, string(char))
	//    }

	p := "💶48.95"
	for i, v := range p {
		fmt.Println(i, v, string(v))
	}
}
