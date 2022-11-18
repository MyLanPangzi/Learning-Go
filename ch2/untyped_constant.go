package main

import "fmt"

func main() {
	const x = 10
	var y int = x
	var z float64 = x
	var d byte = x
	const typedX int = 10
	fmt.Println(x, y, z, d)
	//z = typedX

}
