package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
	if h := t.Hour(); h < 12 {
		fmt.Println("Good morning!")
	} else if h < 17 {
		fmt.Println("Good afternoon.")
	} else {
		fmt.Println("Good evening.")
	}
}

//package main
//
//import (
//	"fmt"
//	"time"
//)
//
//func main() {
//	t := time.Now()
//	switch {
//	case t.Hour() < 12:
//		fmt.Println("Good morning!")
//	case t.Hour() < 17:
//		fmt.Println("Good afternoon.")
//	default:
//		fmt.Println("Good evening.")
//	}
//}
