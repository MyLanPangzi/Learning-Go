package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"os"
)

type LogicProvider struct {
}

func (l LogicProvider) Process(data string) string {
	return fmt.Sprintf("hello %s", data)
}

type Logic interface {
	Process(data string) string
}
type Client struct {
	L Logic
}

func (c Client) Program() {
	data := c.L.Process("world")
	fmt.Println(data)
}

func main() {
	client := Client{
		L: LogicProvider{},
	}
	client.Program()

	//testIoReader()
	testGzipIoReader()
}

func testGzipIoReader() {
	file, err := os.Open("/Users/bo.xie/IdeaProjects/Learning-Go/ch7/interface.go.gz")
	if err != nil {
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	gz, err := gzip.NewReader(file)
	if err != nil {
		return
	}
	defer func(gz *gzip.Reader) {
		err := gz.Close()
		if err != nil {
			panic(err)
		}
	}(gz)

	err = process(gz)
	if err != nil {
		return
	}
}

func testIoReader() {
	file, err := os.Open("/Users/bo.xie/IdeaProjects/Learning-Go/ch7/interface.go")
	if err != nil {
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)
	err = process(file)
	if err != nil {
		return
	}
}

func process(r io.Reader) error {
	bytes, err := io.ReadAll(r)
	if err != nil {
		return err
	}
	fmt.Println(string(bytes))
	return nil
}
