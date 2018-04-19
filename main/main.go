package main

import (
	"fmt"
)

func main() {
	hashMap := map[string]int{
		"a": 1,
		"b": 2,
	}
	for k, v := range hashMap {
		fmt.Println(k, "-", v)
	}
	fmt.Println(hashMap["a"])
}
