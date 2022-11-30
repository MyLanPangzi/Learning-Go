package main

import "fmt"

func main() {
	p := person{}
	i := 1
	s := "hello"
	modifyFails(i, s, p)
	fmt.Println(i, s, p)

	//	 m := map[int]string{
	//        1: "first",
	//        2: "second",
	//    }
	//    modMap(m)
	//    fmt.Println(m)
	//
	//    s := []int{1, 2, 3}
	//    modSlice(s)
	//    fmt.Println(s)

	m := map[int]string{
		1: "1",
		2: "2",
		3: "3",
	}
	modMap(m)
	fmt.Println(m)
	ints := []int{1, 2, 3}
	modSlice(ints)
	fmt.Println(ints)
}

type person struct {
	name string
	age  int
}

func modifyFails(i int, s string, p person) {
	i *= 2
	s = "world"
	p.name = "bob"
}

func modMap(m map[int]string) {
	m[1] = "hello"
	m[2] = "world"
	delete(m, 3)
}

func modSlice(s []int) {
	for i, v := range s {
		s[i] = v * 2
	}
	s = append(s, 10)

}
