package main

import "fmt"

func main() {
	fmt.Println("方法类型的变量")
	a := func() {
		fmt.Println("1")
	}
	a()

	fmt.Println("闭包定义和调用")
	b := func() func() {
		return func() {
			fmt.Println("2")
		}
	}
	bb := b()
	bb()
	b()()

	fmt.Println("闭包")
	c := func() func() {
		return func() {
			fmt.Println("3")
		}
	}()
	c()
}
