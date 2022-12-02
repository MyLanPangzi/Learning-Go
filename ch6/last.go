package main

import (
	"encoding/json"
	"fmt"
)

type Foo struct {
	Name string
	Age  int
}

func main() {
	f := Foo{}
	MakeFoo1(&f)
	fmt.Println(f)

	f2, err := MakeFoo()
	if err != nil {
		panic(err)
	}
	fmt.Println(f2)

	s := struct {
		Name string `json:"name,omitempty"`
		Age  int    `json:"age,omitempty"`
	}{}
	err = json.Unmarshal([]byte(`{"name":"hello", "age": 100}`), &s)
	if err != nil {
		panic(err)
	}
	fmt.Println(s)
}

func MakeFoo() (*Foo, error) {
	return &Foo{
		Name: "world",
		Age:  20,
	}, nil
}

func MakeFoo1(f *Foo) {
	f.Name = "hello"
	f.Age = 10
}
