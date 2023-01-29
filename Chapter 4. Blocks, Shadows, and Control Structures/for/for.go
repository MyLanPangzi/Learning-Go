package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	i := 2
	for i < 10 {
		fmt.Println(i)
		i += 2
	}

	//for {
	//	fmt.Println("hello")
	//	time.Sleep(time.Duration(1000 * 1000 * 1000))
	//}

	for i := 0; i < 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
			continue
		}
		if i%3 == 0 {
			fmt.Println("Fizz")
			continue
		}
		if i%5 == 0 {
			fmt.Println("Buzz")
			continue
		}
	}

	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println(i)
	}
	ints := [...]int{1, 2, 3}
	for i, j := range ints {
		fmt.Println(i, "", j)
	}
	for _, v := range []int{1, 2, 3} {
		fmt.Println(v)
	}
	for k, v := range map[string]bool{
		"hello": true,
		"world": true,
	} {
		fmt.Println(k, v)
	}
	for i, v := range "hello 🌞!" {
		fmt.Println(i, v, string(v))
	}

	vals := []int{1, 2, 3}
	for _, v := range vals {
		v *= 2
	}
	fmt.Println(vals)

	strings := []string{"hello world", "apple"}
outer:
	for _, s := range strings {
		for _, r := range s {
			fmt.Println(string(r))
			if r == 'l' {
				continue outer
			}
		}
	}

	s1 := []int{2, 4, 6, 8, 10}
	for i, v := range s1 {
		if i == 0 {
			continue
		}
		if i == len(s1)-1 {
			continue
		}
		fmt.Println(v)
	}

	for i := 1; i < len(s1)-1; i++ {
		fmt.Println(i, s1[i])
	}
}
