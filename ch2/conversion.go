package main

import "fmt"

func main() {
	var a int = 10
	var b float64 = 20
	var c = int(b) + a
	var d = float64(a) + b
	fmt.Println(a, b, c, d)
}
