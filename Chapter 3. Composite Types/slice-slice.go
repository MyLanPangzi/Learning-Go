package main

import "fmt"

func main() {
	s1 := []int{1, 2, 3, 4, 5, 6}
	s2 := s1[:2]
	s3 := s1[1:]
	s4 := s1[1:3]
	s5 := s1[:]
	fmt.Println(s1, len(s1), cap(s1))
	fmt.Println(s2, len(s2), cap(s2))
	fmt.Println(s3, len(s3), cap(s3))
	fmt.Println(s4, len(s4), cap(s4))
	fmt.Println(s5, len(s5), cap(s5))

	s6 := []int{1, 2, 3, 4}
	s7 := s6[:2]
	s8 := s6[1:]

	s6[1] = 20
	s7[0] = 10
	s8[1] = 30
	fmt.Println(s6)
	fmt.Println(s7)
	fmt.Println(s8)

	s9 := []int{1, 2, 3, 4}
	s10 := s9[:2]
	fmt.Println(cap(s9), cap(s10))
	s10 = append(s10, 20)
	fmt.Println(s9)
	fmt.Println(s10)

}
