package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	if i, err := strconv.ParseInt("100", 2, 8); err == nil {
		fmt.Println(int8(i))
	} else {
		log.Fatal(err)
	}
}
