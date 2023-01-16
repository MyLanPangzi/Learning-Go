package main

import "fmt"

func countTo2(max int) (<-chan int, func()) {
	ch := make(chan int)
	done := make(chan struct{})
	go func() {
		for i := 0; i < max; i++ {
			select {
			case ch <- i:
			case <-done:
				return
			}
		}
		close(ch)
	}()
	return ch, func() {
		close(done)
	}
}
func main() {
	ch, cancel := countTo2(10)
	for v := range ch {
		if v == 5 {
			break
		}
		fmt.Println(v)
	}
	cancel()
}
