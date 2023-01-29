package main

import (
	"fmt"
)

func doPanic(msg string) {
	panic(msg)
}

func div60(i int) {
	defer func() {
		if v := recover(); v != nil {
			//fmt.Println(v)
			fmt.Printf("err : %v\n", v)
		}
	}()
	fmt.Println(60 / i)
}
func main() {
	//doPanic(os.Args[0])
	for _, v := range []int{1, 2, 0, 6} {
		div60(v)
	}
}
