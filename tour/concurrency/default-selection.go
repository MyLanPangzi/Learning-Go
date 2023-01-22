package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(time.Millisecond * 100)
	boom := time.After(time.Millisecond * 500)
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("boom")
			return
		default:
			fmt.Println("     .")
			time.Sleep(time.Millisecond * 50)
		}
	}
}

//func main() {
//	tick := time.Tick(100 * time.Millisecond)
//	boom := time.After(500 * time.Millisecond)
//	for {
//		select {
//		case <-tick:
//			fmt.Println("tick.")
//		case <-boom:
//			fmt.Println("BOOM!")
//			return
//		default:
//			fmt.Println("    .")
//			time.Sleep(50 * time.Millisecond)
//		}
//	}
//}
