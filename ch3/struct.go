package main

import "fmt"

type person struct {
	name string
	age  int
	pet  string
}
type firstPerson struct {
	name string
	age  int
}
type secondPerson struct {
	name string
	age  int
}
type thirdPerson struct {
	age  int
	name string
}
type fourthPerson struct {
	firstName string
	age       int
}

func main() {
	var fred person
	fmt.Println(fred)
	p := person{}
	fmt.Println(p)
	p2 := person{name: "hello", pet: "cat", age: 18}
	fmt.Println(p2)
	p2.name = "world"
	fmt.Println(p2)

	var person2 struct {
		name string
		age  int
		pet  string
	}
	person2.name = "bob"
	person2.age = 50
	person2.pet = "dog"
	fmt.Println(person2)

	pet := struct {
		name string
		kind string
	}{
		name: "fury",
		kind: "dog",
	}
	fmt.Println(pet)

	f1 := firstPerson{}
	s1 := secondPerson{}
	t1 := thirdPerson{}
	four := fourthPerson{}

	s1 = secondPerson(f1)
	f1 = firstPerson(s1)
	//f1 = firstPerson(t1)
	//f1 = firstPerson(four)
	fmt.Println(f1, s1, t1, four)

	g := struct {
		name string
		age  int
	}{
		name: "g",
		age:  10,
	}
	f1 = g
	fmt.Println(g == f1)
}
