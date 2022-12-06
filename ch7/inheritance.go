package main

import "fmt"

type Score int
type Person struct {
	Name string
}
type HighScore Score
type Employee Person

func main() {
	//// assigning untyped constants is valid
	//var i int = 300
	//var s Score = 100
	//var hs HighScore = 200
	//hs = s                  // compilation error!
	//s = i                   // compilation error!
	//s = Score(i)            // ok
	//hs = HighScore(s)       // ok

	var i int = 300
	var s Score = 100
	var hs HighScore = 200
	//hs = s
	//s = i
	s = Score(i)
	hs = HighScore(i)
	hs = HighScore(s)
	fmt.Println(i, s, hs)

}
