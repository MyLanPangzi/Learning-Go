package main

import "fmt"

func main() {
	// for counter := 0; counter < 20; counter++ {
	//        switch(counter / 2) {
	//            case 2, 3, 5, 7:
	//                fmt.Println("Prime value:", counter / 2)
	//            default:
	//                fmt.Println("Non-prime value:", counter / 2)
	//        }
	//    }
	for i := 0; i < 20; i++ {
		switch v := i / 2; v {
		case 2, 3, 5, 7:
			fmt.Println("prime val", v)
		default:
			fmt.Println("non prime", v)
		}
	}
}
