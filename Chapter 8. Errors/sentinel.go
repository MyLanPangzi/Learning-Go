package main

import (
	"archive/zip"
	"bytes"
	"log"
)

func main() {
	data := []byte("This is not a zip file")
	notAZipFile := bytes.NewReader(data)
	_, err := zip.NewReader(notAZipFile, int64(len(data)))
	if err == zip.ErrFormat {
		log.Fatal(err)
		return
	}
}
