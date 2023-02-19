package main

import "fmt"

func main() {
	// products := []string { "Kayak", "Lifejacket", "Paddle", "Hat"}
	//    replacementProducts := []string { "Canoe", "Boots"}
	//    copy(products, replacementProducts)
	//    fmt.Println("products:", products)

	strings := []string{"hello", "world", "go", "java"}
	pr := []string{"py", "rust"}
	copy(strings[0:1], pr)
	//copy(strings, pr)
	fmt.Println(pr, strings)
}
