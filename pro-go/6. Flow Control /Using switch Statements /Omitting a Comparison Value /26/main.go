package main

import "fmt"

func main() {
	// for counter := 0; counter < 10; counter++ {
	//        switch {
	//            case counter == 0:
	//                fmt.Println("Zero value")
	//            case counter < 3:
	//                fmt.Println(counter, "is < 3")
	//            case counter >= 3 && counter < 7:
	//                fmt.Println(counter, "is >= 3 && < 7")
	//            default:
	//                fmt.Println(counter, "is >= 7")
	//        }
	//    }
	for i := 0; i < 10; i++ {
		switch {
		case i == 0:
			fmt.Println("zero")
		case i < 3:
			fmt.Println(i, "<3")
		case i >= 3 && i < 7:
			fmt.Println(i, ">=3 && < 7")
		default:
			fmt.Println(i, ">=", 7)
		}
	}
}
