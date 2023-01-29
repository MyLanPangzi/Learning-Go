package main

import (
	"errors"
	"fmt"
	"os"
	"reflect"
)

func fileChecker2(name string) error {
	f, err := os.Open(name)
	if err != nil {
		return fmt.Errorf("in fileChecker: %w", err)
	}
	f.Close()
	return nil
}

type MyErr struct {
	Codes []int
}

func (me MyErr) GetCodes() []int {
	return me.Codes
}

func (me MyErr) Error() string {
	return fmt.Sprintf("codes : %v", me.Codes)
}

func (me MyErr) Is(target error) bool {
	if me2, ok := target.(MyErr); ok {
		return reflect.DeepEqual(me, me2)
	}
	return false
}

type ResourceErr struct {
	Resource string
	Code     int
}

func (re ResourceErr) Error() string {
	return fmt.Sprintf("%s %d", re.Resource, re.Code)
}

func (re ResourceErr) Is(target error) bool {
	if re2, ok := target.(ResourceErr); ok {
		ignoreRes := re2.Resource == ""
		ignoreCode := re2.Code == 0
		matchRes := re2.Resource == re.Resource
		matchCode := re2.Code == re.Code
		return matchCode && matchRes || matchCode && ignoreRes || matchRes && ignoreCode
	}
	return false
}

func main() {
	err := fileChecker2("not.txt")
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			fmt.Println("file not exist")
		}
	}

	err1 := MyErr{Codes: []int{1, 2, 3}}
	err2 := MyErr{Codes: []int{1, 2, 3}}
	fmt.Println(errors.Is(err1, err2))
	re1 := ResourceErr{
		Resource: "video",
		Code:     400,
	}
	re2 := ResourceErr{
		Resource: "video",
	}
	re3 := ResourceErr{
		Code: 400,
	}
	fmt.Println(errors.Is(re1, re2))
	fmt.Println(errors.Is(re1, re3))

	var myErr MyErr
	var myErr2 error = MyErr{
		Codes: []int{1, 2, 3},
	}
	errors.As(myErr2, &myErr)
	fmt.Println(myErr)
	var myErr3 interface {
		GetCodes() []int
	}
	errors.As(myErr2, &myErr3)
	fmt.Println(myErr3.GetCodes())
}
