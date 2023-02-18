package main

import "fmt"

func main() {
	//kayakPrice := 275.00
	//    if (kayakPrice > 500) {
	//        fmt.Println("Price is greater than 500")
	//    } else if (kayakPrice < 300) {
	//        fmt.Println("Price is less than 300")
	//    }
	p := 60
	if p > 100 {
		fmt.Println("p>100")
	} else if p == 100 {
		fmt.Println("p=100")
	} else if p < 50 {
		fmt.Println("p<50")
	} else {
		fmt.Println(">=50")
	}
}
