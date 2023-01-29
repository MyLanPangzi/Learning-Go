package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	testEmptyInterface()
	testJsonUnmarshal()

	var ll *LinkedList
	ll = ll.Insert(0, 1)
	ll = ll.Insert(0, 2)
	ll = ll.Insert(1, 3)
	fmt.Println(ll, ll.Next, ll.Next.Next)

}

type LinkedList struct {
	Value interface{}
	Next  *LinkedList
}

func (ll *LinkedList) Insert(pos int, val interface{}) *LinkedList {
	if ll == nil || pos == 0 {
		return &LinkedList{
			Value: val,
			Next:  ll,
		}
	}
	ll.Next = ll.Next.Insert(pos-1, val)
	return ll
}

func testEmptyInterface() {
	var i interface{}
	i = 20
	i = "hello"
	i = struct {
		FirstName string
		LastName  string
	}{
		FirstName: "Hello",
		LastName:  "World",
	}
	fmt.Println(i)
}

func testJsonUnmarshal() {
	/*
		data := map[string]interface{}{}
		contents, err := ioutil.ReadFile("testdata/sample.bytes")
		if err != nil {
		    return err
		}
		defer contents.Close()
		bytes.Unmarshal(contents, &data)
	*/
	data := map[string]interface{}{}
	bytes, err := os.ReadFile("/Users/bo.xie/IdeaProjects/Learning-Go/ch7/data.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		panic(err)
	}
	fmt.Println(data)
}
