package main

import "fmt"

func main() {
	//  product := "Kayak"
	//    for index, character := range product {
	//        switch (character) {
	//            case 'K':
	//                fmt.Println("K at position", index)
	//            case 'y':
	//                fmt.Println("y at position", index)
	//        }
	//    }
	s := "hello"
	for i, v := range s {
		switch v {
		case 'h':
			fmt.Println('h', i)
		case 'o':
			fmt.Println('o', i)
		}
	}
}
