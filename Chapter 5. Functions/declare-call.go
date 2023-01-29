package main

import (
	"errors"
	"fmt"
)

func MyFunc(opts Opts) {
	fmt.Println(opts)
}
func main() {
	_, _ = declareAndCall()

	optionalArgs()

	varArgs()

	multipleReturnValue()

}

func declareAndCall() (int, error) {
	return fmt.Println(div(3, 2))
}
func div(numerator, denominator int) int {
	return numerator / denominator
}

func optionalArgs() {
	MyFunc(Opts{})
	MyFunc(Opts{FirstName: "hello"})
	MyFunc(Opts{FirstName: "hello", LastName: "world"})
	MyFunc(Opts{Age: 10})
}

type Opts struct {
	FirstName string
	LastName  string
	Age       int
}

func varArgs() {
	fmt.Println(addTo(10, 1, 2, 3, 4))
	fmt.Println(addTo(10))
	fmt.Println(addTo(10, []int{1, 2, 3}...))
}
func addTo(base int, vals ...int) []int {
	result := make([]int, 0, len(vals))
	for _, val := range vals {
		result = append(result, val+base)
	}
	return result
}

func multipleReturnValue() {
	fmt.Println(divAndRemainder(3, 2))
	fmt.Println(divAndRemainder(3, 0))
	//r, i, err := divAndRemainder(3, 0)
	//if err != nil {
	//	os.Exit(1)
	//}
	//fmt.Println(r, i)
}
func divAndRemainder(numerator, denominator int) (result int, remainder int, err error) {
	result, remainder = 1, 2
	if denominator == 0 {
		return 0, 0, errors.New("cannot div by 0")
	}
	result = numerator / denominator
	remainder = numerator % denominator
	err = nil
	return result, remainder, err
	//return // never do this
}
