package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	c := &SafeCounter{
		m: sync.Mutex{},
		v: map[string]int{},
	}

	for i := 0; i < 1000; i++ {
		go c.Inc("hello")
	}
	time.Sleep(time.Second * 2)
	fmt.Println(c.Value("hello"))
}

type SafeCounter struct {
	m sync.Mutex
	v map[string]int
}

func (c *SafeCounter) Inc(key string) {
	c.m.Lock()
	defer c.m.Unlock()
	c.v[key]++
}
func (c *SafeCounter) Value(key string) int {
	c.m.Lock()
	defer c.m.Unlock()
	return c.v[key]
}
