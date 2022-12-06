package main

import (
	"fmt"
	"io"
)

/*
type MyInt int

	func main() {
	    var i interface{}
	    var mine MyInt = 20
	    i = mine
	    i2 := i.(MyInt)
	    fmt.Println(i2 + 1)
	}
*/
type MyInt int

func main() {
	var i interface{}
	var mine MyInt = 20
	i = mine
	i2 := i.(MyInt)
	fmt.Println(i2)

	i3, ok := i.(int)
	if !ok {
		fmt.Println("type assertion err")
	}
	fmt.Println(i3)
	doThing(1)
	doThing(nil)
	doThing(MyInt(2))
	doThing("hello")
	var r io.Reader
	var ii interface{}
	ii = r
	doThing(r)

	doThing(ii)
	fmt.Println(ii == nil)
	var s *string
	ii = s
	fmt.Println(ii == nil)
	doThing(ii)
	fmt.Println("ii---")
	doThing(false)
	doThing(rune(1))
	doThing(2.0)
}

func doThing(i interface{}) {
	switch j := i.(type) {
	case nil:
		fmt.Println(nil)
	case int:
		fmt.Println("int")
	case MyInt:
		fmt.Println("MyInt")
	case io.Reader:
		fmt.Println("Reader")
	case string:
		fmt.Println("string")
	case bool, rune:
		fmt.Println("bool rune")
	default:
		fmt.Println("nothing", j)
	}
}
