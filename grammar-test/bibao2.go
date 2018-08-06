package main

import "fmt"

func main() {
	a := func(num int) func() {
		return func() {
			fmt.Println("打印外层方法参数:", num)
		}
	}
	a(1)()
	a(2)()
}
