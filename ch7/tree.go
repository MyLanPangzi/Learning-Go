package main

import "fmt"

type IntTree struct {
	val         int
	left, right *IntTree
}

func (it *IntTree) Insert(val int) *IntTree {
	if it == nil {
		return &IntTree{val: val}
	}
	if val < it.val {
		it.left = it.left.Insert(val)
	} else if val > it.val {
		it.right = it.right.Insert(val)
	}
	return it
}
func (it *IntTree) Container(val int) bool {
	switch {
	case it == nil:
		return false
	case val < it.val:
		return it.left.Container(val)
	case val > it.val:
		return it.right.Container(val)
	default:
		return true
	}
}

func main() {
	//var it *IntTree
	//    it = it.Insert(5)
	//    it = it.Insert(3)
	//    it = it.Insert(10)
	//    it = it.Insert(2)
	//    fmt.Println(it.Contains(2))  // true
	//    fmt.Println(it.Contains(12)) // false
	var it *IntTree
	it = it.Insert(5)
	it = it.Insert(3)
	it = it.Insert(10)
	it = it.Insert(2)
	fmt.Println(it.Container(2))
	fmt.Println(it.Container(12))
	fmt.Println(it.Container(10))
	fmt.Println(it)
}
