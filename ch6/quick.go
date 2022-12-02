package main

import "fmt"

func main() {
	//var x int32 = 10
	//var y bool = true
	//pointerX := &x
	//pointerY := &y
	//var pointerZ *string
	var x int32 = 10
	var y bool = true
	px := &x
	py := &y
	var ps *string
	fmt.Println(*px, *py, *px+10, ps, ps == nil)
	//fmt.Println(*ps)

	s := "hello"
	ps = &s
	fmt.Println(ps, *ps)

	i := new(int)
	fmt.Println(i, i == nil, *i)
	*i = 100
	fmt.Println(i, *i)
	type Foo struct {
		Name string
	}
	f := &Foo{Name: "hello"}
	fmt.Println(f, *f)

	type person struct {
		FirstName  string
		MiddleName *string
		LastName   string
	}
	stringp := func(s string) *string {
		return &s
	}
	p := person{
		FirstName:  "hello",
		MiddleName: stringp("!"),
		LastName:   "world",
	}
	fmt.Println(p)

}
