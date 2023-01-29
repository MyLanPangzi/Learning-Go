package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	t, err := time.Parse("2006-02-01 15:04:05 -0700", "2016-13-03 00:00:00 +0000")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(t)
	fmt.Println(t.Format("2006-01-02 15:04:05 -0700"))
}
