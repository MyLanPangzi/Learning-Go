package main

import "fmt"

type I3 interface {
	M()
}
type T3 struct {
	S string
}

func (t *T3) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}
func desc3(i I3) {
	fmt.Printf("%v %T\n", i, i)
}
func main() {
	var i I3
	fmt.Println(i == nil)
	var t *T3
	i = t
	fmt.Println(i == nil)
	desc3(i)
	i.M()

	i = &T3{"hello"}
	desc3(i)
	i.M()
}

//func main() {
//	var i I
//
//	var t *T
//	i = t
//	describe(i)
//	i.M()
//
//	i = &T{"hello"}
//	describe(i)
//	i.M()
//}
//
//func describe(i I) {
//	fmt.Printf("(%v, %T)\n", i, i)
//}
