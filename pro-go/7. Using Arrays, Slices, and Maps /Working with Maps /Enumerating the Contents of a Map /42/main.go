package main

import "fmt"

func main() {
	//  products := map[string]float64 {
	//        "Kayak" : 279,
	//        "Lifejacket": 48.95,
	//        "Hat": 0,
	//    }
	//    for key, value := range products {
	//        fmt.Println("Key:", key, "Value:", value)
	//    }

	m := map[string]int{
		"hello": 10,
		"world": 10,
		"go":    10,
	}
	for k, v := range m {
		fmt.Println(k, v)
	}
}
