package main

import "fmt"

func main() {
	first := 100
	//const second = 200.01 Invalid operation: first == second (cannot convert the constant second to the type int)
	const second = 200.0
	fmt.Println(first == second, first != second, first > second, first >= second, first < second, first <= second)

}
