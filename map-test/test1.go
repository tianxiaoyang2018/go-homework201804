package main

import "fmt"

type Bean struct {
	A int
	B string
}

func main() {
	map1 := make(map[Bean]string)
	map1[Bean{A: 1, B: "1"}] = "abc"
	map1[Bean{A: 1, B: "1"}] = "edf"
	map1[Bean{A: 2, B: "2"}] = "ghi"
	for k, v := range map1 {
		fmt.Println(k, v)
	}
}
