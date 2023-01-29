package main

import "fmt"

type Status int

const (
	Invalidation Status = iota + 1
	NotFound
)

type StatusErr struct {
	Status  Status
	Message string
}

func (se StatusErr) Error() string {
	return se.Message
}

//func GenerateError(flag bool) error {
//    var genErr StatusErr
//    if flag {
//        genErr = StatusErr{
//            Status: NotFound,
//        }
//    }
//    return genErr
//}

func GeneratorError(flag bool) error {
	var genErr StatusErr
	if flag {
		genErr = StatusErr{
			Status: NotFound,
		}
	}
	return genErr
}
func GeneratorError2(flag bool) error {
	if flag {
		return StatusErr{Status: NotFound}
	}
	return nil
}
func GeneratorError3(flag bool) error {
	var genErr error
	if flag {
		return StatusErr{Status: NotFound}
	}
	return genErr
}
func main() {

	err := GeneratorError(true)
	fmt.Println(err != nil)
	err = GeneratorError(false)
	fmt.Println(err != nil)

	err = GeneratorError2(true)
	fmt.Println(err != nil)
	err = GeneratorError2(false)
	fmt.Println(err != nil)

	err = GeneratorError3(true)
	fmt.Println(err != nil)
	err = GeneratorError3(false)
	fmt.Println(err != nil)
}
