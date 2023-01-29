package main

import (
	"fmt"
	"strconv"
)

func main() {
	var s = "hello world"
	b := s[6]
	fmt.Println(s, b)
	fmt.Println(s[4:7], s[:5], s[6:])

	var s2 = "Hello 🌞"
	fmt.Println(s2[4:7], s2[:5], s2[6:])
	fmt.Println(len(s2))

	var a rune = 'x'
	var c byte = 'y'
	fmt.Println(string(a), string(c), string(65), string(strconv.Itoa(65)))
	fmt.Println([]byte(s2), []rune(s2))
	fmt.Println("a" > "b", "a" == "a")
}
