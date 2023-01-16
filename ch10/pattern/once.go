package main

import (
	"fmt"
	"sync"
	"time"
)

type SlowComplicatedParser interface {
	Parse(string) string
}

var parser SlowComplicatedParser
var once sync.Once

func Parse(data string) string {
	once.Do(func() {
		parser = initParser()
	})
	return parser.Parse(data)
}

func initParser() SlowComplicatedParser {
	return MyParser{}
}

type MyParser struct {
}

func (m MyParser) Parse(s string) string {
	time.Sleep(time.Millisecond * 100)
	return fmt.Sprintf("hello %v", s)
}

func main() {
	fmt.Println(Parse("world"))
}
