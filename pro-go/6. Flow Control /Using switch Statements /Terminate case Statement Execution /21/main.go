package main

import "fmt"

func main() {
	// product := "Kayak"
	//    for index, character := range product {
	//        switch (character) {
	//            case 'K', 'k':
	//                if (character == 'k') {
	//                    fmt.Println("Lowercase k at position", index)
	//                    break
	//                }
	//                fmt.Println("Uppercase K at position", index)
	//            case 'y':
	//                fmt.Println("y at position", index)
	//        }
	//    }

	s := "hello Hello"
	for i, c := range s {
		switch c {
		case 'h', 'H':
			if c == 'h' {
				fmt.Println("h at ", i)
				break
			}
			fmt.Println("H at ", i)
		case 'o':
			fmt.Println("o at ", i)
		}
	}
}
