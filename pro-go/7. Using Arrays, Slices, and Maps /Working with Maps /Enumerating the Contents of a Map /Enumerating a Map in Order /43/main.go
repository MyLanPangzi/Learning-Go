package main

import (
	"fmt"
	"sort"
)

func main() {
	//  products := map[string]float64 {
	//        "Kayak" : 279,
	//        "Lifejacket": 48.95,
	//        "Hat": 0,
	//    }
	//    keys := make([]string, 0, len(products))
	//    for key, _ := range products {
	//        keys = append(keys, key)
	//    }
	//    sort.Strings(keys)
	//    for _, key := range keys {
	//        fmt.Println("Key:", key, "Value:", products[key])
	//    }

	m := map[string]int{
		"c": 10,
		"a": 10,
		"b": 10,
	}
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		fmt.Println(k, m[k])
	}
}
