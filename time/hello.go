package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	month := now.Month()

	fmt.Println(now.Weekday(), month)
}
