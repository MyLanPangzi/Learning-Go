package main

import "fmt"

func main() {
	var price float32 = 2
	var tax float32 = 0.2
	fmt.Println(price + tax)
	price = 300
	fmt.Println(price + tax)

}
