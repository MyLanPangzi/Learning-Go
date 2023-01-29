package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"unicode"
)

func main() {
	s := "The quick brown fox jumped over the lazy dog"
	reader := strings.NewReader(s)
	m, err := countLetters(reader)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(m)

	gZipReader, closer, err := buildGZipReader("reader.go.gz")
	if err != nil {
		log.Fatal(err)
	}
	defer closer()
	m, err = countLetters(gZipReader)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(m)
}

func countLetters(r io.Reader) (map[string]int, error) {
	m := map[string]int{}
	buf := make([]byte, 2048)
	for {
		n, err := r.Read(buf)
		for _, b := range buf[:n] {
			if unicode.IsLetter(rune(b)) {
				m[string(b)]++
			}
		}
		if err == io.EOF {
			return m, nil
		}
		if err != nil {
			return nil, err
		}
	}
}
func buildGZipReader(fileName string) (*gzip.Reader, func(), error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, nil, err
	}
	gr, err := gzip.NewReader(file)
	if err != nil {
		return nil, nil, err
	}
	return gr, func() {
		gr.Close()
		file.Close()
	}, nil
}
