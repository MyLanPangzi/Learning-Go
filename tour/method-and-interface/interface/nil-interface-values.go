package main

import "fmt"

func main() {
	var i I4
	desc4(i)
	i.M()
}

type I4 interface {
	M()
}

func desc4(i I4) {
	fmt.Printf("%v %T\n", i, i)
}

//func main() {
//	var i I
//	describe(i)
//	i.M()
//}
//
//func describe(i I) {
//	fmt.Printf("(%v, %T)\n", i, i)
//}
