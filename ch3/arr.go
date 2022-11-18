package main

import "fmt"

func main() {
	var x [3]int
	fmt.Println(x)

	var arr = [3]int{10, 20, 30}
	fmt.Println(arr)

	var arr2 = [12]int{1, 5: 4, 6, 10: 100, 15}
	fmt.Println(arr2)

	var arr3 = [...]int{1, 2, 3}
	fmt.Println(arr3)

	var arr4 = [3]int{1, 2, 3}
	fmt.Println(arr4 == arr3)

	var arr5 [2][3]int
	fmt.Println(arr5)

	var a6 [1]int
	a6[0] = 10
	fmt.Println(a6)
	fmt.Println(len(a6))
}
