package main

import "fmt"

func main() {
	s := []string{"1", "2"}
	for i, v := range s {
		fmt.Println(i, v)
	}
}
