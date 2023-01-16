package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func processAndGather(in <-chan int, processor func(int) int, num int) []int {
	out := make(chan int, num)
	wg := sync.WaitGroup{}
	wg.Add(num)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < num; i++ {
		go func() {
			defer wg.Done()
			for v := range in {
				time.Sleep(time.Millisecond * time.Duration(rand.Int63n(1000)))
				out <- processor(v)
			}
		}()
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	var result []int
	for v := range out {
		result = append(result, v)
	}
	return result
}
func main() {
	//testWg()
	in := make(chan int, 10)
	go func() {
		for i := 0; i < 10; i++ {
			in <- i
			time.Sleep(time.Millisecond * 100)
		}
		close(in)
	}()

	r := processAndGather(in, func(i int) int {
		return i * 2
	}, 10)
	fmt.Println(r)
}

func testWg() {
	wg := sync.WaitGroup{}
	wg.Add(3)

	for i := 0; i < 3; i++ {
		go func() {
			defer wg.Done()
			time.Sleep(time.Second)
		}()
	}
	wg.Wait()
	fmt.Println("done")
}
