package main

import "fmt"

func main() {
	//c := make(chan int)
	//done := make(chan int)
	//go func() {
	//	for i := 0; i < 10; i++ {
	//		fmt.Println(<-c)
	//	}
	//	close(done)
	//}()
	//f(c, done)
	ch1 := make(chan int)
	ch2 := make(chan int)
	select {
	case <-ch1:
	case <-ch2:
		//default:
		//	fmt.Println("default")
	}

}
func f(c, done chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-done:
			fmt.Println()
			return
		}
	}
}

//func main() {
//	c := make(chan int)
//	quit := make(chan int)
//	go func() {
//		for i := 0; i < 10; i++ {
//			fmt.Println(<-c)
//		}
//		quit <- 0
//	}()
//	fibonacci(c, quit)
//}
