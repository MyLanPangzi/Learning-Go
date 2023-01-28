package main

import (
	"fmt"
)

func processChannel(ch chan int) []int {
	const conc = 10
	results := make(chan int, conc)
	for i := 0; i < conc; i++ {
		go func() {
			v := <-ch
			results <- process(v)
		}()
	}
	var out []int
	for i := 0; i < conc; i++ {
		out = append(out, <-results)
	}
	return out
}

func process(v int) int {
	return v * 2
}
func main() {
	in := make(chan int, 10)
	for i := 0; i < 10; i++ {
		in <- i
	}
	for _, v := range processChannel(in) {
		fmt.Println(v)
	}
}
