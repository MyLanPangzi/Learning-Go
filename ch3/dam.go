package main

import "fmt"

func main() {
	x := make([]int, 0, 5)
	x = append(x, 1, 2, 3, 4)
	y := x[:2]
	z := x[2:]
	//y := x[:2:2]
	//z := x[2:4:4]
	fmt.Println(cap(x), cap(y), cap(z))
	fmt.Println(len(x), len(y), len(z))
	y = append(y, 30, 40, 50)
	x = append(x, 60)
	z = append(z, 70)
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("z:", z, len(z), cap(z))

	s11 := []int{1, 2, 3, 4}
	s12 := s11[:2:2]
	s13 := s11[2:4:4]
	s12 = append(s12, 5)
	s13 = append(s13, 6)
	fmt.Println(s11)
	fmt.Println(s12)
	fmt.Println(s13)
}
