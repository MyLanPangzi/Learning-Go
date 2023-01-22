package main

import "fmt"

func main() {
	var l *List[int]
	head := l.add(1)
	l = l.add(2)
	fmt.Println(l, head)
}

type List[T any] struct {
	next *List[T]
	val  T
}

func (l *List[T]) add(e T) *List[T] {
	if l == nil {
		return &List[T]{
			val: e,
		}
	}
	return l.next.add(e)
}

//type List[T any] struct {
//	next *List[T]
//	val  T
//}
//
//func main() {
//}
