package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	//var price []rune = []rune("€48.95")
	//    var currency string = string(price[0])
	//    var amountString string = string(price[1:])
	//    amount, parseErr  := strconv.ParseFloat(amountString, 64)
	//    fmt.Println("Length:", len(price))
	//    fmt.Println("Currency:", currency)
	//    if (parseErr == nil) {
	//        fmt.Println("Amount:", amount)
	//    } else {
	//        fmt.Println("Parse Error:", parseErr)
	//    }
	p := []rune("€48.95")
	c := string(p[0])
	a := string(p[1:])
	if f, err := strconv.ParseFloat(a, 64); err == nil {
		fmt.Println(c, f, len(p))
	} else {
		log.Fatal(err)
	}
}
