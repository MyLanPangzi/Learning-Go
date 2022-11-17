package main

import "fmt"

func main() {
	var s string = "hello world"
	var t string = `hello 
world`
	fmt.Println(s, t)
}
