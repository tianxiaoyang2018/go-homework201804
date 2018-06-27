package main

import "fmt"

func main() {
	add := func(a int) func(int) int {
		return func(b int)int {
			return (a+b)
		}
	}
	fmt.Println(add(1)(2))
	fmt.Println(add(2)(3))

}
