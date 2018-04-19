package main

import "fmt"

type Phone interface {
	call()
}

type NokiaPhone struct {
}

func (nokiaPhone NokiaPhone) call() {
	fmt.Println("诺基亚")
}

type IPhone struct {
}

func (iPhone IPhone) call() {
	fmt.Println("苹果")
}

func test(phone Phone) {
	phone.call()
}

func main() {
	var phone Phone

	phone = new(NokiaPhone)
	phone.call()

	test(new(IPhone))
}
