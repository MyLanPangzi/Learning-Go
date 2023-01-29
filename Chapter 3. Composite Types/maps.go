package main

import "fmt"

func main() {
	var nilMap map[string]int
	fmt.Println(nilMap, nilMap == nil)
	fmt.Println(map[string]int{} == nil)

	var team = map[string][]string{
		"a": {"hello"},
		"b": {"world"},
		"c": {"!"},
	}
	fmt.Println(team)

	m := make(map[int]string, 10)
	fmt.Println(m, len(m))

	m2 := map[string]int{}
	m2["hello"] = 1
	fmt.Println(m2["hello"], m2["world"])
	m2["world"]++
	fmt.Println(m2["world"])
	v, ok := m2["go"]
	fmt.Println(v, ok)
	v, ok = m2["hello"]
	fmt.Println(v, ok)
	delete(m2, "hello")
	//delete(m2, "hello")
	v, ok = m2["hello"]
	fmt.Println(v, ok)

	var m3 = map[int]bool{}
	ints := []int{1, 2, 3, 4, 1, 2, 3, 4}
	for _, v := range ints {
		m3[v] = true
	}
	fmt.Println(m3, len(m3), m3[5])

	var m4 = map[int]struct{}{}
	for _, v := range ints {
		m4[v] = struct{}{}
	}
	s, ok := m4[5]
	fmt.Println(m4, len(m4), s, ok)
}
