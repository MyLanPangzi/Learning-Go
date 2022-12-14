package main

import (
	"fmt"
	"sort"
)

func main() {
	funcAsParam()
	funcAsReturnVal()
}

func funcAsParam() {

	type Person struct {
		FirstName string
		LastName  string
		Age       int
	}
	people := []Person{
		{"Pat", "Patterson", 37},
		{"Tracy", "Bobbert", 23},
		{"Fred", "Fredson", 18},
	}
	fmt.Println(people)

	sort.Slice(people, func(i, j int) bool {
		return people[i].LastName < people[j].LastName
	})
	fmt.Println(people)

	sort.Slice(people, func(i, j int) bool {
		return people[i].Age < people[j].Age
	})
	fmt.Println(people)

}

func funcAsReturnVal() {
	fmt.Println(makeMult(10)(20))
	//	    twoBase := makeMult(2)
	//    threeBase := makeMult(3)
	//    for i := 0; i < 3; i++ {
	//        fmt.Println(twoBase(i), threeBase(i))
	//    }
	//}
	twoBase := makeMult(2)
	threeBase := makeMult(3)
	for i := 0; i < 3; i++ {
		fmt.Println(twoBase(i), threeBase(i))
	}
}

func makeMult(base int) func(int) int {
	return func(factor int) int {
		return factor * base
	}
}
