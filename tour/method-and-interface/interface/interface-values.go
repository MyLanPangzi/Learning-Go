package main

import (
	"fmt"
	"math"
)

type I2 interface {
	M()
}
type T2 struct {
	S string
}

func (t T2) M() {
	fmt.Println(t.S)
}

type F float64

func (f F) M() {
	fmt.Println(f)
}

func main() {
	var i I2
	var i2 I2
	fmt.Println(i == i2, i == nil)

	i = &T2{"hello"}
	desc(i)
	i.M()

	i2 = F(math.Pi)
	desc(i2)
	i2.M()

	//i = i2

	fmt.Println(i == i2)
}
func desc(i I2) {
	fmt.Printf("%v %T\n", i, i)
}

//func main() {
//	var i I
//
//	i = &T{"Hello"}
//	describe(i)
//	i.M()
//
//	i = F(math.Pi)
//	describe(i)
//	i.M()
//}
//
//func describe(i I) {
//	fmt.Printf("(%v, %T)\n", i, i)
//}
