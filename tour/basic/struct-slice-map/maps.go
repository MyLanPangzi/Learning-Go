package main

import "fmt"

func main() {
	type Vertex struct {
		Lat, Long float64
	}
	var m map[string]Vertex
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		Lat:  40.68433,
		Long: -74.39967,
	}
	fmt.Println(m)
}

//package main
//
//import "fmt"
//
//type Vertex struct {
//	Lat, Long float64
//}
//
//var m map[string]Vertex
//
//func main() {
//	m = make(map[string]Vertex)
//	m["Bell Labs"] = Vertex{
//		40.68433, -74.39967,
//	}
//	fmt.Println(m["Bell Labs"])
//}
