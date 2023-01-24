package main

import "sync"

var mutex = sync.Mutex{}

func main() {
	a()
}
func a() {
	mutex.Lock()
	defer mutex.Unlock()
	b()
}

func b() {
	mutex.Lock()
	defer mutex.Unlock()
}
