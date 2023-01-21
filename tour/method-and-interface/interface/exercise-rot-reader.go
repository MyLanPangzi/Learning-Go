package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	_, err := io.Copy(os.Stdout, r)
	if err != nil {
		panic(err)
	}
	fmt.Println()

	s = strings.NewReader("Lbh penpxrq gur pbqr!")
	builder := &strings.Builder{}
	r = rot13Reader{s}
	io.Copy(builder, r)
	fmt.Println(builder.String())

	s = strings.NewReader(builder.String())
	r = rot13Reader{s}
	io.Copy(os.Stdout, r)
	fmt.Println()
}

type rot13Reader struct {
	r io.Reader
}

func (r rot13Reader) Read(b []byte) (n int, err error) {
	n, err = r.r.Read(b)
	for i := range b {
		b[i] = rot13(b[i])
	}
	return n, err
}
func rot13(x byte) byte {
	capital := x >= 'A' && x <= 'Z'
	if !capital && (x < 'a' || x > 'z') {
		return x // Not a letter
	}

	x += 13
	if capital && x > 'Z' || !capital && x > 'z' {
		x -= 26
	}
	return x
}

//type rot13Reader struct {
//	r io.Reader
//}
//
//func main() {
//	s := strings.NewReader("Lbh penpxrq gur pbqr!")
//	r := rot13Reader{s}
//	io.Copy(os.Stdout, &r)
//}
