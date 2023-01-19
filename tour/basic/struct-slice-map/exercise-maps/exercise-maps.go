package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

//Implement WordCount. It should return a map of the counts of each “word” in the string s. The wc.Test function runs a test suite against the provided function and prints success or failure.
//
//You might find strings.Fields helpful.

func main() {
	wc.Test(WordCount)
}

func WordCount(s string) map[string]int {
	fields := strings.Fields(s)
	m := make(map[string]int)
	for _, v := range fields {
		m[v]++
	}
	return m
}

//package main
//
//import (
//	"golang.org/x/tour/wc"
//)
//
//func WordCount(s string) map[string]int {
//	return map[string]int{"x": 1}
//}
//
//func main() {
//	wc.Test(WordCount)
//}
