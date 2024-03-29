package main

import "fmt"

func main() {
	type Vertex struct {
		Lat, Long float64
	}
	m := map[string]Vertex{
		"Bell Labs": {40, -74},
		"Google":    {37, -122},
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
//var m = map[string]Vertex{
//	"Bell Labs": {40.68433, -74.39967},
//	"Google":    {37.42202, -122.08408},
//}
//
//func main() {
//	fmt.Println(m)
//}
