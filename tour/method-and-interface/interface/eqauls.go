package main

import (
	"fmt"
)

func main() {
	var i1 I8 = T8{"hello"}
	t := T8{"hello"}
	var i2 I8 = t
	fmt.Println(i1 == i2)
	i3 := i2
	fmt.Println(i3 == i2)
	t.S = "world"
	//v := reflect.ValueOf(i2)
	//name := v.FieldByName("S")
	//fmt.Println(name.CanSet())
	fmt.Println(i3 == i2, i2, i3)

}

type I8 interface {
	M()
}
type T8 struct {
	S string
}

func (t T8) M() {
	fmt.Println(t.S)
}
