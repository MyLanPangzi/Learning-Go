package main

import (
	"fmt"
	"math/rand"
	"time"
)

// package main
//
// import (
//
//	"fmt"
//	"math/rand"
//
// )
//
//	func main() {
//		fmt.Println("My favorite number is", rand.Intn(10))
//	}
func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("My favorite number is ", rand.Intn(10))
}
