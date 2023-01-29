package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func searchData(s string, searchers []func(string) []string) []string {
	done := make(chan struct{})
	result := make(chan []string)
	for _, searcher := range searchers {
		go func(searcher func(string) []string) {
			select {
			case result <- searcher(s):
			case <-done:

			}
		}(searcher)
	}
	r := <-result
	close(done)
	return r
}

func main() {
	rand.Seed(time.Now().UnixNano())
	var searchers []func(string) []string
	for i := 0; i < 10; i++ {
		searchers = append(searchers, func(s string) []string {
			time.Sleep(time.Duration(int(time.Millisecond) * rand.Intn(1000)))
			n := rand.Intn(100)
			return []string{strconv.Itoa(n), strings.Repeat(s, n)}
		})
	}
	data := searchData("h", searchers)
	fmt.Println(data)

}
