package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

func main() {
	type t struct {
		Name string `json:"name,omitempty"`
		Age  int    `json:"age,omitempty"`
	}
	data := `{"name": "Fred", "age": 40}
{"name": "Mary", "age": 21}
{"name": "Pat", "age": 30}`
	decoder := json.NewDecoder(strings.NewReader(data))
	var ts []t
	for decoder.More() {
		var temp t
		err := decoder.Decode(&temp)
		ts = append(ts, temp)
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Printf("%+v\n", ts)

	var b bytes.Buffer
	builder := &strings.Builder{}
	encoder := json.NewEncoder(builder)
	encoder2 := json.NewEncoder(&b)
	for _, v := range ts {
		err := encoder.Encode(v)
		if err != nil {
			log.Fatal(err)
		}
		err = encoder2.Encode(v)
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println(builder.String())
	fmt.Println(b.String())
}
