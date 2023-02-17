package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	if i, err := strconv.Atoi("100"); err == nil {
		fmt.Println(i)
	} else {
		log.Fatal(err)
	}
}
