package main

import (
	"fmt"
	"strconv"
)

func main() {
	//i := ImpossibleStruct[int]{10} //int does not implement ImpossiblePrintableInt (missing String method)
	//fmt.Println(i)
	t := ImpossibleStruct[MyInt]{10} // MyInt does not implement ImpossiblePrintableInt (possibly missing ~ for int in constraint ImpossiblePrintableInt)
	fmt.Println(t)
}

type ImpossiblePrintableInt interface {
	//int
	~int
	fmt.Stringer
}
type ImpossibleStruct[T ImpossiblePrintableInt] struct {
	v T
}
type MyInt int

func (i MyInt) String() string {
	return strconv.Itoa(int(i))
}
