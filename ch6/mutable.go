package main

import "fmt"

func main() {
	var f *int
	failedUpdate(f)
	fmt.Println(f)

	x := 10
	failedUpdate(&x)
	fmt.Println(x)
	update(&x)
	fmt.Println(x)
}

func failedUpdate(f *int) {
	x := 10
	f = &x
}

func update(i *int) {
	*i = 20
}
