package main

import "fmt"

type MailCategory int

//const (
//    Uncategorized MailCategory = iota
//    Personal
//    Spam
//    Social
//    Advertisements
//)

const (
	Uncategorized MailCategory = iota
	Personal
	Spam
	Social
	Advertisements
)

// type BitField int
//
// const (
//
//	Field1 BitField = 1 << iota // assigned 1
//	Field2                      // assigned 2
//	Field3                      // assigned 4
//	Field4                      // assigned 8
//
// )
type BitField int

const (
	F1 BitField = 1 << iota
	F2
	F3
	F4
)

func main() {
	fmt.Println(Uncategorized)
	fmt.Println(Personal)
	fmt.Println(Spam)
	fmt.Println(Social)
	fmt.Println(Advertisements)

	fmt.Println(0 << 1)
	fmt.Println(1 << 1)
	fmt.Println(2 << 1)
}
