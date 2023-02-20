package main

import "fmt"

func main() {
	//	 products := make(map[string]float64, 10)
	//    products["Kayak"] = 279
	//    products["Lifejacket"] = 48.95
	//    fmt.Println("Map size:", len(products))
	//    fmt.Println("Price:", products["Kayak"])
	//    fmt.Println("Price:", products["Hat"])
	m := make(map[string]float64, 10)
	m["hello"] = 10
	fmt.Println(len(m), m["hello"], m["world"], m)
}
