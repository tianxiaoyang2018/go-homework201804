package main

import "fmt"

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	arr := []int{1,2,3}
	fmt.Printf("我擦%+v",arr)
}
