package main

import (
	"errors"
	"fmt"
	"log"
	"strconv"
)

func DoSomeThings(val1 int, val2 string) (_ string, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("in DoSomeThings: %w", err)
		}
	}()
	val3, err := doThing1(val1)
	if err != nil {
		return "", err
		//return "", fmt.Errorf("in DoSomeThings: %w", err)
	}
	val4, err := doThing2(val2)
	if err != nil {
		return "", err
		//return "", fmt.Errorf("in DoSomeThings: %w", err)
	}
	r, err := doThing3(val3, val4)
	if err != nil {
		return "", err
		//return "", fmt.Errorf("in DoSomeThings: %w", err)
	}
	return r, nil
}

func doThing3(val3 string, val4 string) (string, error) {
	return val3 + ", " + val4, nil
}

func doThing2(val2 string) (string, error) {
	return val2, nil
}

func doThing1(val1 int) (string, error) {
	if val1 == 1 {
		return "", errors.New("val1 is cannot be 1")
	}
	return strconv.Itoa(val1), nil
}
func main() {
	things, err := DoSomeThings(1, "world")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(things)

}
