package main

import "fmt"

const Pi = 3.14

func main() {
	const World = "世界"
	const (
		b = true
		i = 10
	)
	fmt.Println(b, i)
	fmt.Println("Hello ", World)
	fmt.Println(Pi)
}

//package main
//
//import "fmt"
//
//const Pi = 3.14
//
//func main() {
//	const World = "世界"
//	fmt.Println("Hello", World)
//	fmt.Println("Happy", Pi, "Day")
//
//	const Truth = true
//	fmt.Println("Go rules?", Truth)
//}
