package main

import "golang.org/x/tour/reader"

type MyReader struct {
}

//Implement a Reader type that emits an infinite stream of the ASCII character 'A'.

func (m MyReader) Read(p []byte) (n int, err error) {
	p[0] = 'A'
	return 1, nil
}

func main() {
	reader.Validate(MyReader{})
}

//package main
//
//import "golang.org/x/tour/reader"
//
//type MyReader struct{}
//
//// TODO: Add a Read([]byte) (int, error) method to MyReader.
//
//func main() {
//	reader.Validate(MyReader{})
//}
