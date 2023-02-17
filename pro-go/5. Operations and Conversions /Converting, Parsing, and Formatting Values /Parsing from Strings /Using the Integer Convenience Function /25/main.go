package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	if i, err := strconv.ParseInt("100", 10, 0); err == nil {
		fmt.Println(int(i))
	} else {
		log.Fatal(err)
	}
}
