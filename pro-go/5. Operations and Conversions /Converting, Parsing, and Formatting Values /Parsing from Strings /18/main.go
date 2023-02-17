package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	val1 := "0"
	b, err := strconv.ParseBool(val1)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(b)
	}

}
