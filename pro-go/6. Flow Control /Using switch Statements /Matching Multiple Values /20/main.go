package main

import "fmt"

func main() {
	// product := "Kayak"
	//    for index, character := range product {
	//        switch (character) {
	//            case 'K', 'k':
	//                fmt.Println("K or k at position", index)
	//            case 'y':
	//                fmt.Println("y at position", index)
	//        }
	//    }
	s := "hello Hello"
	for i, c := range s {
		switch c {
		case 'h', 'H':
			fmt.Println("h or H", i)
		case 'o':
			fmt.Println("o", i)

		}
	}
}
