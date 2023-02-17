package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	if i, err := strconv.ParseInt("100", 0, 8); err == nil {
		fmt.Println(i)
	} else {
		log.Fatal(err)
	}
}
