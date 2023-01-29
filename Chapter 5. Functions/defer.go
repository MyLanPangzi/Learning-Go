package main

import (
	"io"
	"log"
	"os"
)

func main() {
	readFile()
}

func readFile() {
	if len(os.Args) < 2 {
		log.Fatalln("no file specified")
	}

	f, closer, err := getFile(os.Args[1])
	if err != nil {
		return
	}
	defer closer()
	//f, err := os.Open(os.Args[1])
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//
	//defer func(f *os.File) {
	//	_ = f.Close()
	//}(f)

	in := make([]byte, 2048)
	for {
		count, err := f.Read(in)
		if err != nil {
			if err != io.EOF {
				log.Fatalln(err)
				return
			}
			return
		}
		count, err = os.Stdout.Write(in[:count])
		if err != nil {
			log.Fatalln(err)
			return
		}
	}
}

func getFile(fileName string) (*os.File, func(), error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, nil, err
	}
	return file, func() {
		_ = file.Close()
	}, nil
}
