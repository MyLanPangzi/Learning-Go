package main

import "fmt"

func main() {
	fmt.Println("couting")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i) //Possible resource leak, 'defer' is called in the 'for' loop
	}
	fmt.Println("done")
}

//package main
//
//import "fmt"
//
//func main() {
//	fmt.Println("counting")
//
//	for i := 0; i < 10; i++ {
//		defer fmt.Println(i)
//	}
//
//	fmt.Println("done")
//}
