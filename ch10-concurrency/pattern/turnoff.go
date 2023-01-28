package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	done := make(chan struct{})
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		wg.Wait()
		close(done)
	}()
	go func() {
		for i := 0; i < 2; i++ {
			ch1 <- i
		}
		close(ch1)
		wg.Done()
	}()
	go func() {
		for i := 0; i < 3; i++ {
			ch2 <- i
		}
		close(ch2)
		wg.Done()
	}()
	for {
		select {
		case i, ok := <-ch1:
			if !ok {
				ch1 = nil
				continue
			}
			fmt.Println("ch1", i)
		case i, ok := <-ch2:
			if !ok {
				ch2 = nil
				continue
			}
			fmt.Println("ch2", i)
		case <-done:
			goto alldone
		default:
			fmt.Println("nothing")
			time.Sleep(time.Millisecond * 500)
		}
	}
alldone:
	fmt.Println("done")
}
