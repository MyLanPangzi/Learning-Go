package main

import "fmt"

func main() {
	arr := [4]int{1, 2, 3, 4}
	a1 := arr[:2]
	a2 := arr[2:]
	fmt.Println(arr, a1, a2)

	a1[0] = 0
	fmt.Println(a1, arr)
	fmt.Println(len(a1), cap(a1))
}
