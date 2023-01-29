package main

import (
	"fmt"
	"strconv"
)

func main() {
	funcAsValue()
	anonymousFunc()
}

type opFuncType func(int, int) int

func funcAsValue() {
	ops := map[string]opFuncType{
		"+": add,
		"-": sub,
		"*": mul,
		"/": divide,
	}
	fmt.Println(ops["+"](1, 2))
	fmt.Println(ops["-"](1, 2))
	fmt.Println(ops["*"](1, 2))
	fmt.Println(ops["/"](4, 2))

	expressions := [][]string{
		{"2", "+", "1"},
		{"2", "-", "1"},
		{"2", "*", "1"},
		{"2", "/", "1"},
		{"2", "%", "1"},
		{"two", "+", "three"},
		{"3"},
	}
	for _, expression := range expressions {
		if len(expression) < 3 {
			fmt.Println("invalid expression", expression)
			continue
		}

		p1, err := strconv.Atoi(expression[0])
		if err != nil {
			fmt.Println(err)
			continue
		}

		op := expression[1]
		opFunc, ok := ops[op]
		if !ok {
			fmt.Println("unsupported op", op)
			continue
		}

		p2, err := strconv.Atoi(expression[0])
		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Println(opFunc(p1, p2))
	}
}

func anonymousFunc() {
	for i := 0; i < 3; i++ {
		func(j int) {
			fmt.Println(j)
		}(i)
	}
}

func add(i, j int) int {
	return i + j
}
func sub(i, j int) int {
	return i - j
}
func divide(i, j int) int {
	return i / j
}
func mul(i, j int) int {
	return i * j
}
