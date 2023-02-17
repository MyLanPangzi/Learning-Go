package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	if b, err := strconv.ParseBool("0"); err == nil {
		fmt.Println(b)
	} else {
		log.Fatal(err)
	}
}
