package main

import (
	"fmt"
	"time"
)

func main() {
	f()
}

func f() {
	aStream := make(chan string)
	bStream := make(chan string)
	cStream := make(chan string)
	done := make(chan interface{})
	time.AfterFunc(time.Millisecond*5000, func() {
		close(done)
	})

	go func() {
		defer close(aStream)
		for {
			select {
			case <-done:
				return
			default:
				aStream <- "A"
			}
		}

	}()
	go func() {
		defer close(bStream)
		for {
			select {
			case <-done:
				return
			default:

			}
			bStream <- "B"

		}
	}()
	go func() {
		defer close(cStream)
		for {
			select {
			case <-done:
				return
			default:
				cStream <- "C"
			}
		}
	}()
	go func() {
		for {
			select {
			case <-done:
				return
			case a := <-aStream:
				fmt.Println(a)
				select {
				case <-done:
					return
				case b := <-bStream:
					fmt.Println(b)
					select {
					case <-done:
						return
					case c := <-cStream:
						fmt.Println(c)

					}
				}
				time.Sleep(1000 * time.Millisecond)
			}
		}
	}()
	<-done
}
