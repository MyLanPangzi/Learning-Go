package main

import (
	"fmt"
	"time"
)

func main() {
	var or func(channels ...<-chan int) <-chan int
	// main -> or(2h,1h,5m,1m,1s,1ms,1us):or1 <-or1
	//go1 select 2h,1h,5m,or(1m,1s,1ms,1us,or1):or2 defer close(or1)
	//go2 select 1m,1s,1ms,or(1us,or1,or2):or3 defer close(or2)
	//go3 select 1us,or1,or2,or(or3):or3 defer close(or3)
	or = func(channels ...<-chan int) <-chan int {
		switch len(channels) {
		case 0:
			return nil
		case 1:
			return channels[0]
		}

		orDone := make(chan int)
		go func() {
			defer close(orDone)

			switch len(channels) {
			case 2:
				select {
				case <-channels[0]:
				case <-channels[1]:
				}
			default:
				select {
				case <-channels[0]:
				case <-channels[1]:
				case <-channels[2]:
				case <-or(append(channels[3:], orDone)...):
				}
			}
		}()
		return orDone
	}
	sig := func(after time.Duration) <-chan int {
		c := make(chan int)
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	start := time.Now()
	<-or(
		sig(2*time.Hour),
		sig(1*time.Hour),
		sig(5*time.Minute),
		sig(1*time.Minute),
		sig(1*time.Second),
		sig(1*time.Millisecond),
		sig(1*time.Microsecond),
	)

	fmt.Printf("done after %v\n", time.Since(start))
}
