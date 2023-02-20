package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	//var price string = "$48.95"
	//    var currency byte = price[0]
	//    var amountString string = price[1:]
	//    amount, parseErr  := strconv.ParseFloat(amountString, 64)
	//    fmt.Println("Currency:", currency)
	//    if (parseErr == nil) {
	//        fmt.Println("Amount:", amount)
	//    } else {
	//        fmt.Println("Parse Error:", parseErr)
	//    }
	p := "$48.95"
	c := p[0]
	fmt.Println("currency", c, string(c))
	if a, err := strconv.ParseFloat(p[1:], 64); err == nil {
		fmt.Println("amount", a)
	} else {
		panic(err)
	}

	p = "€48.95"
	c = p[0]
	fmt.Println(len(p))
	fmt.Println("currency", c, string(c))
	if a, err := strconv.ParseFloat(p[1:], 64); err == nil {
		fmt.Println("amount", a)
	} else {
		log.Fatal(err)
	}

}
