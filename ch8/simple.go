package main

import (
	"fmt"
	"log"
)

func doubleEven(i int) (int, error) {
	if i%2 != 0 {
		return 0, fmt.Errorf("%d isn't an even number", i)
		//return 0, errors.New("only even numbers are processed")
	}
	return i * 2, nil
}
func main() {
	_, err := doubleEven(2)
	if err != nil {
		log.Fatal(err)
	}
	double, err := doubleEven(3)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(double)
}
