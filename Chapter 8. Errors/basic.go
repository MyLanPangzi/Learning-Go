package main

import (
	"errors"
	"fmt"
	"log"
)

func calcRemainderAndMod(numerator, denominator int) (int, int, error) {
	if denominator == 0 {
		return 0, 0, errors.New("denominator is 0")
	}
	return numerator / denominator, numerator % denominator, nil

}
func main() {
	numerator := 20
	denominator := 3
	remainder, mod, err := calcRemainderAndMod(numerator, denominator)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(remainder, mod)
}
