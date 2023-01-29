package main

import "fmt"

func main() {
	var i int = 10
	var j = 20
	fmt.Println(i, j)

	var k, l int
	fmt.Println(k, l)

	var x, y = 1, 2
	fmt.Println(x, y)

	var s1, a = "hello", 1
	fmt.Println(s1, a)

	var (
		a1     int
		a2     int = 10
		a3         = "hello"
		a4, a5     = "world", 'a'
	)
	fmt.Println(a1, a2, a3, a4, a5)

	z := 100
	fmt.Println(z)

	q, w := 10, 20
	fmt.Println(q, w)
}
