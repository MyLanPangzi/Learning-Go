package main

import "fmt"

func main() {
	words := []string{"a", "cow", "smile", "gopher", "octopus", "1234567890"}
	for _, word := range words {
		switch size := len(word); size {
		case 1, 2, 3, 4:
			fmt.Println(word, "short")
		case 5:
			fmt.Println(word, "good")
		case 6, 7, 8, 9:
		default:
			fmt.Println(word, "long")
		}
	}

loop:
	for i := 0; i < 10; i++ {
		switch {
		case i%2 == 0:
			fmt.Println(i, "even")
		case i%3 == 0:
			fmt.Println(i, "mod 3 is 0")
		case i%7 == 0:
			fmt.Println("seven")
			break loop
		default:
			fmt.Println(i, "boring")
		}
	}
	strings := []string{"hello", "world", "go", "spark", "java", "flink", "streaming", "dolphinsscheduler"}
	for _, s := range strings {
		switch l := len(s); {
		case l < 5:
			fmt.Println(s, "short")
		case l > 10:
			fmt.Println(s, "long")
		default:
			fmt.Println(s, "good")
		}
	}
	a := 0
	switch {
	case a == 1:
		fmt.Println(1)
	case a == 2:
		fmt.Println(2)
	default:
		fmt.Println(a)
	}

	switch a {
	case 1:
		fmt.Println(1)
	default:
		fmt.Println(a)
	}
}
