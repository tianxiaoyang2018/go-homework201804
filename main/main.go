package main

import "fmt"

func main() {
	arr := []string{"abc","def"}
	for _, str := range arr {
		str = "我"
		fmt.Println(str)
	}
	fmt.Println(arr)
}
