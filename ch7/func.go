package main

import "fmt"

type Adder struct {
	start int
}

func (a Adder) AddTo(val int) int {
	return a.start + val
}

func main() {
	adder := Adder{start: 10}
	fmt.Println(adder.AddTo(5))
	f := adder.AddTo
	fmt.Println(f(20))
	f2 := Adder.AddTo
	fmt.Println(f2(adder, 30))

}
