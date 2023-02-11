package main

import "fmt"

func DoubleToString(i int) string {
	return fmt.Sprintf("%d", i*2)
}
func main() {
	s := []int{10, 20, 30}
	strings := Map(s, DoubleToString)
	for _, v := range strings {
		fmt.Println(v)
	}
	sum := Reduce(s, 0, func(x int, y int) int {
		return x + y
	})
	fmt.Println(sum)

	filter := Filter(s, func(i int) bool {
		return i > 10
	})
	for _, v := range filter {
		fmt.Println(v)
	}
}
func Map[T1, T2 any](s []T1, f func(T1) T2) []T2 {
	t2 := make([]T2, len(s))
	for i, v := range s {
		t2[i] = f(v)
	}
	return t2
}
func Reduce[T1, T2 any](s []T1, initializer T2, f func(T2, T1) T2) T2 {
	r := initializer
	for _, v := range s {
		r = f(r, v)
	}
	return r

}
func Filter[T any](s []T, f func(T) bool) []T {
	r := make([]T, 0, len(s))
	for _, v := range s {
		if f(v) {
			r = append(r, v)
		}
	}
	return r
}
