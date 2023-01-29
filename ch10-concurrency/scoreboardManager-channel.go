package main

import (
	"fmt"
	"sync"
)

func main() {
	manager, cancel := NewChannelScoreboardManager()
	defer cancel()
	manager.Update("hello", 1)
	fmt.Println(manager.Read("hello"))

	var wg sync.WaitGroup
	wg.Add(20)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			manager.Inc("hello")
		}()
	}
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			manager.Inc2("hello")
		}()
	}
	wg.Wait()
	fmt.Println(manager.Read("hello"))
}

type ChannelScoreboardManager chan func(map[string]int)

func NewChannelScoreboardManager() (ChannelScoreboardManager, func()) {
	ch := make(ChannelScoreboardManager)
	done := make(chan struct{})
	go func() {
		scoreboard := map[string]int{}
		for {
			select {
			case <-done:
				return
			case f := <-ch:
				f(scoreboard)
			}
		}
	}()

	return ch, func() {
		close(done)
	}
}

func (csm ChannelScoreboardManager) Inc(name string) {
	csm <- func(m map[string]int) {
		m[name]++
	}
}
func (csm ChannelScoreboardManager) Inc2(name string) {
	csm <- func(m map[string]int) {
		m[name]++
	}
}
func (csm ChannelScoreboardManager) Update(name string, val int) {
	csm <- func(m map[string]int) {
		m[name] = val
	}
}
func (csm ChannelScoreboardManager) Read(name string) (int, bool) {
	var out int
	var ok bool
	var wg sync.WaitGroup
	wg.Add(1)
	csm <- func(m map[string]int) {
		defer wg.Done()
		out, ok = m[name]
	}
	wg.Wait()
	return out, ok
}
