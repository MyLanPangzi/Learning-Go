package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {

	fmt.Println(parseBool("true"), parseBool("false"))
	b, err := strconv.ParseBool("not true")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(b)
}

func parseBool(s string) bool {
	b, err := strconv.ParseBool(s)
	if err != nil {
		panic(err)
	}
	return b
}
