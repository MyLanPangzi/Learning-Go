package main

import (
	"fmt"
	"time"
)

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v , %s", e.When, e.What)
}
func run() error {
	return &MyError{
		When: time.Now(),
		What: "it didn't work",
	}
}

//func main() {
//	if err := run(); err != nil {
//		fmt.Println(err)
//	}
//}
