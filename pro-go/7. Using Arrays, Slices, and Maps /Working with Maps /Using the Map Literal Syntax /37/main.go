package main

import "fmt"

func main() {
	//  products := map[string]float64 {
	//        "Kayak" : 279,
	//        "Lifejacket": 48.95,
	//    }
	//    fmt.Println("Map size:", len(products))
	//    fmt.Println("Price:", products["Kayak"])
	//    fmt.Println("Price:", products["Hat"])
	m := map[string]float64{
		"hello": 10,
	}
	fmt.Println(len(m), m, m["hello"], m["world"])
}
