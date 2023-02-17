package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	if i, err := strconv.ParseInt("100", 0, 8); err == nil {
		i8 := int8(i)
		fmt.Printf("%T %v\n", i8, i8)
	} else {
		log.Fatalf("%T %v\n", err, err)
	}

}
