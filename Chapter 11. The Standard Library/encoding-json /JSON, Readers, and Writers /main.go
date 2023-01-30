package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Person struct {
	Name string `json:"name,omitempty"`
	Age  int    `json:"age,omitempty"`
}

func main() {
	toFile := Person{
		Name: "hello",
		Age:  19,
	}
	temp, err := os.CreateTemp(os.TempDir(), "sample-")
	if err != nil {
		return
	}
	defer os.Remove(temp.Name())
	fmt.Println(temp.Name())
	err = json.NewEncoder(temp).Encode(toFile)
	if err != nil {
		return
	}
	temp.Close()

	temp2, err := os.Open(temp.Name())
	if err != nil {
		return
	}
	defer temp2.Close()
	var fromFile Person
	err = json.NewDecoder(temp2).Decode(&fromFile)
	if err != nil {
		return
	}
	fmt.Printf("%+v\n", fromFile)
}
