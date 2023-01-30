package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type Item struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Order struct {
	Id          string    `json:"id"`
	DateOrdered time.Time `json:"date_ordered"`
	CustomerId  string    `json:"customer_id"`
	Items       []Item    `json:"items"`
}

func main() {
	bytes, err := os.ReadFile("/Users/bo.xie/IdeaProjects/Learning-Go/Chapter 11. The Standard Library/encoding-json /data.json")
	if err != nil {
		return
	}
	var order Order
	err = json.Unmarshal(bytes, &order)
	if err != nil {
		return
	}
	fmt.Println(order)
	s, err := json.Marshal(order)
	if err != nil {
		return
	}
	fmt.Println(string(s))
}
