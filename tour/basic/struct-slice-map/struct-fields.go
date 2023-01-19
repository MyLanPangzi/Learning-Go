package main

import "fmt"

func main() {
	type Vertex struct {
		X int
		Y int
	}
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)
}

//package main
//
//import "fmt"
//
//type Vertex struct {
//	X int
//	Y int
//}
//
//func main() {
//	v := Vertex{1, 2}
//	v.X = 4
//	fmt.Println(v.X)
//}
