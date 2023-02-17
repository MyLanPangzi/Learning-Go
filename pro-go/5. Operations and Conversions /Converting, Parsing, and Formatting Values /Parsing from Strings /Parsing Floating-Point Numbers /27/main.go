package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	if f, err := strconv.ParseFloat("2.0", 64); err == nil {
		fmt.Println(f)
	} else {
		log.Fatal(err)
	}
}
