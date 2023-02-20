package main

import "fmt"

func main() {
	// products := map[string]float64 {
	//        "Kayak" : 279,
	//        "Lifejacket": 48.95,
	//        "Hat": 0,
	//    }
	//    value, ok := products["Hat"]
	//    if (ok) {
	//        fmt.Println("Stored value:", value)
	//    } else {
	//        fmt.Println("No stored value")
	//    }
	m := map[string]float64{
		"hello": 10,
	}
	s, ok := m["hello"]
	fmt.Println(s, ok)
	s, ok = m["world"]
	fmt.Println(s, ok)
	if s, ok := m["go"]; ok {
		fmt.Println(s)
	}
}
