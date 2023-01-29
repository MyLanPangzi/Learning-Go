package main

import "fmt"

/*
var s *string
fmt.Println(s == nil) // prints true
var i interface{}
fmt.Println(i == nil) // prints true
i = s
fmt.Println(i == nil) // prints false
*/
func main() {
	var s *string
	fmt.Println(s == nil)
	var i interface{}
	fmt.Println(i == nil)
	i = s
	fmt.Println(i == nil)
}
