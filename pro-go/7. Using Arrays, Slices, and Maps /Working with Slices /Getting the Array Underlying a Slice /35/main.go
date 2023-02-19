package main

import "fmt"

func main() {
	//  p1 := []string { "Kayak", "Lifejacket", "Paddle", "Hat"}
	//    arrayPtr := (*[3]string)(p1)
	//    array := *arrayPtr
	//    fmt.Println(array)

	s := []string{"1", "2", "3"}
	arrPtr := (*[2]string)(s)
	for i, v := range arrPtr {
		fmt.Println(i, v)
	}
	for i, v := range *arrPtr {
		fmt.Println(i, v)
	}
}
