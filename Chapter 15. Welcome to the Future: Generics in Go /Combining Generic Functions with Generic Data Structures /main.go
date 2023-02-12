package main

import (
	"fmt"
	"strings"
)

type BuiltInOrdered interface {
	~string | ~int | ~uint | ~uintptr |
		~int8 | ~int16 | ~int32 | ~int64 |
		~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float64 | ~float32
}

func BuiltInOrder[T BuiltInOrdered](x, y T) int {
	if x < y {
		return -1
	}
	if x > y {
		return 1
	}
	return 0
}

type OrderableFunc[T any] func(x, y T) int

type Node[T any] struct {
	val         T
	left, right *Node[T]
}
type Tree[T any] struct {
	f    OrderableFunc[T]
	root *Node[T]
}

func NewTree[T any](f OrderableFunc[T]) *Tree[T] {
	return &Tree[T]{f: f}
}
func (t *Tree[T]) PreTraverse(f func(T)) {
	if t == nil {
		return
	}
	if t.root == nil {
		return
	}
	t.root.PreTraverse(f)
}
func (t *Tree[T]) Add(v T) {
	t.root = t.root.Add(t.f, v)
}
func (t *Tree[T]) Contains(v T) bool {
	return t.root.Contains(t.f, v)
}

func (n *Node[T]) PreTraverse(f func(T)) {
	if n == nil {
		return
	}
	n.left.PreTraverse(f)
	f(n.val)
	n.right.PreTraverse(f)
}

func (n *Node[T]) Add(f OrderableFunc[T], v T) *Node[T] {
	if n == nil {
		return &Node[T]{val: v}
	}
	switch r := f(v, n.val); {
	case r < 0:
		n.left = n.left.Add(f, v)
	case r > 0:
		n.right = n.right.Add(f, v)
	}
	return n

}
func (n *Node[T]) Contains(f OrderableFunc[T], v T) bool {
	if n == nil {
		return false
	}
	switch r := f(n.val, v); {
	case r < 0:
		return n.left.Contains(f, v)
	case r > 0:
		return n.right.Contains(f, v)
	}
	return true

}

type Person struct {
	Name string
	Age  int
}

func (p Person) Order(other Person) int {
	r := strings.Compare(p.Name, other.Name)
	if r == 0 {
		return p.Age - other.Age
	}
	return r
}
func OrderPeople(x, y Person) int {
	f := Person.Order
	return f(x, y)
}

func main() {
	t1 := NewTree(BuiltInOrder[int])
	t1.Add(10)
	t1.Add(30)
	t1.Add(15)
	t1.Add(1)
	fmt.Println(t1.Contains(10), t1.Contains(20), t1.Contains(15))
	fmt.Println(t1.root.val, t1.root.right.val, t1.root.right.left.val)
	t1.PreTraverse(func(i int) {
		fmt.Println(i)
	})
	t2 := NewTree(OrderPeople)
	t2.Add(Person{"Bob", 30})
	t2.Add(Person{"Maria", 35})
	t2.Add(Person{"Bob", 50})
	t2.Add(Person{"Alice", 50})
	fmt.Println(t2.Contains(Person{"Bob", 30}), t2.Contains(Person{"Bob", 90}))
	t2.PreTraverse(func(person Person) {
		fmt.Println(person)
	})

	//	t3 := NewTree(Person.Order)
	//	t3.Add(Person{"Bob", 30})
	//	t3.Add(Person{"Maria", 35})
	//	t3.Add(Person{"Bob", 50})
	//	fmt.Println(t3.Contains(Person{"Bob", 30}))
	//	fmt.Println(t3.Contains(Person{"Fred", 25}))
	//}
	//
}
