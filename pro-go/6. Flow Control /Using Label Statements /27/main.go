package main

import "fmt"

func main() {
	//counter := 0
	//    target: fmt.Println("Counter", counter)
	//    counter++
	//    if (counter < 5) {
	//        goto target
	//    }
	i := 0
target:
	fmt.Println("i", i)
	i++
	if i < 5 {
		goto target
	}
}
