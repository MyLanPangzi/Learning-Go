package main

import (
	"errors"
	"fmt"
	"time"
)

func timeLimit() (int, error) {
	var result int
	var err error
	done := make(chan struct{})
	go func() {
		result, err = do()
		close(done)
	}()
	select {
	case <-done:
		return result, err
	case <-time.After(time.Second * 2):
		return 0, errors.New("time out")
	}
}

func do() (int, error) {
	//time.Sleep(time.Millisecond * 2500)
	time.Sleep(time.Millisecond * 500)
	//time.Sleep(time.Second * 63 << 2)
	return 1, nil
}
func main() {
	fmt.Println(timeLimit())
}
