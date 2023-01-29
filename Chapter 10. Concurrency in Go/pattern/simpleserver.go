package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	client := &http.Client{
		Timeout: time.Second * 10,
	}
	for i := 0; i < 100; i++ {
		go request(client)
	}
	time.Sleep(time.Second * 3)
}

func request(client *http.Client) {
	response, err := client.Get("http://localhost:8080/request")
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	data := make([]byte, response.ContentLength)
	response.Body.Read(data)
	fmt.Println(string(data))
}
