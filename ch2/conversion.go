package main

import "fmt"

func main() {
	//var a int = 10
	//var b float64 = 20
	//var c = int(b) + a
	//var d = float64(a) + b
	//fmt.Println(a, b, c, d)

	var x int = 10
	var y float64 = 30.2
	var z float64 = float64(x) + y
	var d int = x + int(y)
	fmt.Println(z, d)
}
