package main

import "fmt"

func main() {
	var s Stack[int]
	for i := 0; i < 100; i++ {
		s.Push(i)
	}
	fmt.Println(s.Contains(1), s.Contains(101))
	//for i := 0; i < 101; i++ {
	//	fmt.Println(s.Pop())
	//}
}

type Stack[T comparable] struct {
	vals []T
}

func (s *Stack[T]) Push(val T) {
	s.vals = append(s.vals, val)
}
func (s *Stack[T]) Pop() (T, bool) {

	if len(s.vals) == 0 {
		var zero T
		return zero, false
	}
	top := s.vals[len(s.vals)-1]
	s.vals = s.vals[:len(s.vals)-1]
	return top, true
}
func (s *Stack[T]) Contains(val T) bool {
	for _, v := range s.vals {
		if v == val {
			return true
		}
	}
	return false
}
