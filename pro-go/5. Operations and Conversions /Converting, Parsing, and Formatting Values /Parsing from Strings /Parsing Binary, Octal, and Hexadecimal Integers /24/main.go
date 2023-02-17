package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	//if i, err := strconv.ParseInt("0B1100100", 0, 8); err == nil {
	if i, err := strconv.ParseInt("0b1100100", 0, 8); err == nil {
		fmt.Println(int8(i))
	} else {
		log.Fatal(err)
	}
}
