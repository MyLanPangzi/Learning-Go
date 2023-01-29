package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	f, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			panic(err)
		}
	}(f)

	bytes := make([]byte, 100)
	for {
		count, err := f.Read(bytes)
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		fmt.Print(string(bytes[:count]))
	}
}
