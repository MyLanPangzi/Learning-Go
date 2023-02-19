package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	// products := []string { "Kayak", "Lifejacket", "Paddle", "Hat"}
	//    sort.Strings(products)
	//    for index, value := range products {
	//        fmt.Println("Index:", index, "Value:", value)
	//    }
	s := []string{"3", "1", "2"}
	sort.Strings(s)
	fmt.Println(s)
	ints := []int{3, 1, 2}
	sort.Ints(ints)
	fmt.Println(ints)
	float64s := []float64{3, 2, 1}
	sort.Float64s(float64s)
	fmt.Println(float64s)

	people := []Person{
		{"hello", 1},
		{"world", 2},
		{"go", 3},
	}
	sort.Slice(people, func(i, j int) bool {
		return strings.Compare(people[i].Name, people[j].Name) < 0
	})
	fmt.Println(people)

	var names Names = []string{"3", "2", "1"}
	sort.Sort(names)
	fmt.Println(names)
}

type Names []string

func (n Names) Len() int {
	return len(n)
}

func (n Names) Less(i, j int) bool {
	return strings.Compare(n[i], n[j]) < 0
}

func (n Names) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}

type Person struct {
	Name string
	Age  int
}
