package main

import "fmt"

func main() {
	//products := map[string]float64 {
	//        "Kayak" : 279,
	//        "Lifejacket": 48.95,
	//        "Hat": 0,
	//    }
	//    delete(products, "Hat")
	//    if value, ok := products["Hat"]; ok {
	//        fmt.Println("Stored value:", value)
	//    } else {
	//        fmt.Println("No stored value")
	//    }

	m := map[string]int{
		"hello": 10,
		"world": 10,
	}
	delete(m, "world")
	if s, ok := m["world"]; ok {
		fmt.Println(s)
	}
}
