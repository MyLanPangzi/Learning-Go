package main

import "fmt"

func main() {
	// kayakPrice := 275.00
	//    if (kayakPrice > 500) {
	//        scopedVar := 500
	//        fmt.Println("Price is greater than", scopedVar)
	//    } else if (kayakPrice < 100) {
	//        scopedVar := "Price is less than 100"
	//        fmt.Println(scopedVar)
	//    } else {
	//        scopedVar := false
	//        fmt.Println("Matched: ", scopedVar)
	//    }

	if p := 275.0; p > 500 {
		i := 500
		fmt.Println("p > ", i)
	} else if p < 100 {
		i := "p<100"
		fmt.Println(i)
	} else {
		i := false
		fmt.Println("matched", i)
	}
}
