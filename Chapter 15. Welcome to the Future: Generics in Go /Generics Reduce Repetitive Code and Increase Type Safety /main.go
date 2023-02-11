package main

import "fmt"

type Node struct {
	val  int
	next *Node
}
type Orderable interface {
	Order(val any) int
}
type Tree struct {
	val   Orderable
	left  *Tree
	right *Tree
}

func (t *Tree) Insert(val Orderable) *Tree {
	if t == nil {
		return &Tree{val: val}
	}
	switch cmp := val.Order(t.val); {
	case cmp < 0:
		t.left = t.left.Insert(val)
	case cmp > 0:
		t.left = t.left.Insert(val)
	}
	return t
}

type OrderInt int

func (i OrderInt) Order(val any) int {
	return int(i - val.(OrderInt))
}
func main() {
	var t *Tree
	t = t.Insert(OrderInt(1))
	t = t.Insert(OrderInt(2))
	t = t.Insert(OrderInt(3))
	fmt.Println(t)
}

func Min(x, y float64) float64 {
	if x < y {
		return x
	}
	return y
}
