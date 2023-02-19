package main

import "fmt"

func main() {
	//  products := [4]string { "Kayak", "Lifejacket", "Paddle", "Hat"}
	//    deleted := append(products[:2], products[3:]...)
	//    fmt.Println("Deleted:", deleted)
	nums := []string{"1", "2", "3", "4"}
	nums = append(nums[:2], nums[3:]...)
	//nums.remove
	fmt.Println(nums, len(nums), cap(nums))
}
