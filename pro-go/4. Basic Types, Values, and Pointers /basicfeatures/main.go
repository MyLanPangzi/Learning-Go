package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println(rand.Int(), rand.Intn(10))

	var r = 'A'
	fmt.Println(r)
	fmt.Println(10, 10.0, "hello", true, 'A')
	fmt.Println(reflect.TypeOf(10), reflect.TypeOf(10.0),
		reflect.TypeOf("hello"), reflect.TypeOf(true), reflect.TypeOf('A'))

	const price float32 = 2.0
	const tax float32 = 0.2
	const quantity = 2
	fmt.Println(price+tax, (price+tax)*quantity)
}
