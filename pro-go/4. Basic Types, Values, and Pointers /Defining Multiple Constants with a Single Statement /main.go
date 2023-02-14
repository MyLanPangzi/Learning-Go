package main

import "fmt"

func main() {
	const price, tax float32 = 2, 0.2
	const qantity, inStock = 2, true
	fmt.Println("total: ", 2*qantity*(price+tax))
	fmt.Println("In stock:", inStock)
}
