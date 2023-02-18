package main

import "fmt"

func main() {
	//product := "Kayak"
	//    for index, character := range product {
	//        switch (character) {
	//            case 'K':
	//                fmt.Println("Uppercase character")
	//                fallthrough
	//            case 'k':
	//                fmt.Println("k at position", index)
	//            case 'y':
	//                fmt.Println("y at position", index)
	//        }
	//    }
	s := "hello Hello"
	for i, c := range s {
		switch c {
		case 'H':
			fmt.Println("Upper H")
			fallthrough
		case 'h':
			fmt.Println("h at ", i)
		case 'o':
			fmt.Println("o at ", i)

		}
	}
}
