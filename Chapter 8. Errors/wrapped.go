package main

import (
	"errors"
	"fmt"
	"os"
)

func fileChecker(name string) error {
	f, err := os.Open(name)
	if err != nil {
		return fmt.Errorf("in fileChecker: %w", err)
	}
	f.Close()
	return nil
}
func main() {
	err := fileChecker("not.txt")
	if err != nil {
		fmt.Println(err)
		if wrappedErr := errors.Unwrap(err); wrappedErr != nil {
			fmt.Println(wrappedErr)
		}
	}

	err2 := WrapErr{
		Err:     errors.New("wrapped err"),
		Message: "outer err",
	}
	fmt.Println(err2)
	fmt.Println(errors.Unwrap(err2))
}

type WrapErr struct {
	Err     error
	Message string
}

func (w WrapErr) Error() string {
	return w.Message
}

func (w WrapErr) Unwrap() error {
	return w.Err
}
