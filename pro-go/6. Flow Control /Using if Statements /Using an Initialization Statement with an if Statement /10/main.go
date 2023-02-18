package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	// priceString := "275"
	//    if kayakPrice, err := strconv.Atoi(priceString); err == nil {
	//        fmt.Println("Price:", kayakPrice)
	//    } else {
	//        fmt.Println("Error:", err)
	//    }

	if p, err := strconv.Atoi("275"); err == nil {
		fmt.Println("p:", p)
	} else {
		log.Fatal(err)
	}
}
