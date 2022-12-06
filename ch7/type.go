package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func (p Person) String() string {
	return fmt.Sprintf("%s %s age %d", p.FirstName, p.LastName, p.Age)
}

type Score int
type Convert func(string) Score
type TeamScores map[string]Score

// return fmt.Sprintf("%s %s, age %d", p.FirstName, p.LastName, p.Age)
func main() {
	person := Person{
		FirstName: "Hello",
		LastName:  "World",
		Age:       18,
	}
	fmt.Println(person)
}
