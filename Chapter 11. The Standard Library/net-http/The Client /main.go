package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	client := &http.Client{
		Timeout: 30 * time.Second,
	}
	url := "https://jsonplaceholder.typicode.com/todos/1"
	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("hello", "world")
	response, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		log.Fatal(response.Status)
	}
	var data Data
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", data)
}

type Data struct {
	UserId    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}
