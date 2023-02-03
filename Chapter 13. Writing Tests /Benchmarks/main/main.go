package main

import (
	"os"
	"strings"
)

func main() {
	file := "/Users/bo.xie/IdeaProjects/Learning-Go/Chapter 13. Writing Tests /Benchmarks/bench/testdata/data.txt"

	repeat := strings.Repeat("1", 65204)
	err := os.WriteFile(file, []byte(repeat), os.ModeAppend)
	if err != nil {
		return
	}
}
