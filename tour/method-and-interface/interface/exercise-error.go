package main

import "fmt"

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	z := 1.0
	pre := 0.0
	for i := 0; i < 10 && pre != z; i++ {
		pre = z
		z -= (z*z - x) / (2 * z)
		fmt.Println(pre, z)
	}
	return z, nil
}
