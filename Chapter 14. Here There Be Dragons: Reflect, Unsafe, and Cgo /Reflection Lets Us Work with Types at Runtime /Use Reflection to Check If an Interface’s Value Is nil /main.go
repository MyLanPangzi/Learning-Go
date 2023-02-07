package main

import (
	"fmt"
	"nil-interface/check"
)

func main() {
	var i any
	fmt.Println(check.HasNoValue(i), i == nil)
	var p *int
	i = p
	fmt.Println(check.HasNoValue(i), i == nil)
}
