package main

import (
	"fmt"
	"time"
)

type Counter struct {
	total       int
	lastUpdated time.Time
}

func (c Counter) String() string {
	return fmt.Sprintf("total %d last update ", c.total, c.lastUpdated)
}

func (c *Counter) Increment() {
	c.total++
	c.lastUpdated = time.Now()
}

func doUpdateWrong(c Counter) {
	c.Increment()
	fmt.Println("wrong ", c.String())
}
func doUpdateRight(c *Counter) {
	c.Increment()
	fmt.Println("right ", c.String())
}

func main() {
	counter := Counter{
		total:       0,
		lastUpdated: time.Now(),
	}
	fmt.Println(counter)
	counter.Increment()
	fmt.Println(counter)

}
