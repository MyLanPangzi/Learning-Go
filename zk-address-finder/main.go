package main

import (
	"encoding/xml"
	"fmt"
	"log"
	"os"
)

type Configuration struct {
	Properties []Property `xml:"property"`
}
type Property struct {
	Name        string `xml:"name"`
	Value       string `xml:"value"`
	Description string `xml:"description"`
}

func main() {
	bytes, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	c := Configuration{}
	err = xml.Unmarshal(bytes, &c)
	if err != nil {
		log.Fatal(err)
	}
	for _, prop := range c.Properties {
		if prop.Name == "yarn.resourcemanager.zk-address" {
			fmt.Println(prop.Value)
		}
	}
}
