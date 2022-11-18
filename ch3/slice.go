package main

import "fmt"

func main() {
	ints := []int{1, 2, 3}
	fmt.Println(ints)

	var s1 = []int{1, 2, 3}
	fmt.Println(s1)
	//fmt.Println(s1==ints)
	fmt.Println(s1 == nil)

	var s3 = []int{1, 5: 4, 6, 10: 100, 15}
	fmt.Println(s3)

	var s4 [][]int
	fmt.Println(s4, len(s4), cap(s4))

	//var s5 []int
	//s5[0] = 1
	//fmt.Println(s5)

	var s6 = []int{1}
	s6[0] = 10
	fmt.Println(s6)

	var s7 []int
	fmt.Println(s7 == nil)

	//var arr [0]int
	//fmt.Println(arr==nil)

	var s8 []int
	s8 = append(s8, 1)
	fmt.Println(s8)
	s8 = append(s8, 2, 3, 4)
	fmt.Println(s8)

	var s9 []int
	s9 = append(s9, s8...)
	fmt.Println(s9)

	var s10 []int
	fmt.Println(s10, len(s10), cap(s10))
	s10 = append(s10, 1)
	fmt.Println(s10, len(s10), cap(s10))
	s10 = append(s10, 2)
	fmt.Println(s10, len(s10), cap(s10))
	s10 = append(s10, 3)
	fmt.Println(s10, len(s10), cap(s10))
	s10 = append(s10, 4)
	fmt.Println(s10, len(s10), cap(s10))
	s10 = append(s10, 5)
	fmt.Println(s10, len(s10), cap(s10))

	var s11 = make([]int, 5)
	fmt.Println(s11, len(s11), cap(s11))

	s12 := make([]int, 5)
	s12 = append(s12, 6)
	fmt.Println(s12, len(s12), cap(s12))

	s13 := make([]int, 5, 10)
	fmt.Println(s13, len(s13), cap(s13))

	s14 := make([]int, 0, 10)
	fmt.Println(s14, len(s14), cap(s14))
	s14 = append(s14, 1, 2, 3, 4, 5)
	fmt.Println(s14, len(s14), cap(s14))

	var s15 []int
	fmt.Println(s15, len(s15), cap(s15), s15 == nil)

	//goland:noinspection GoPreferNilSlice
	var s16 = []int{}
	fmt.Println(s16, len(s16), cap(s16), s16 == nil)

	s17 := []int{1, 1, 1}
	fmt.Println(s17, len(s17), cap(s17))
}
