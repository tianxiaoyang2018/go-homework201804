package main

import "fmt"

var ctest2  chan int
var strtest2 string

func f() {
	strtest2 = "123"
	ctest2<-1
}

func main() {
	ctest2 = make(chan int,0)
	strtest2 = "abc"
	go f()
	<-ctest2
	fmt.Println(strtest2)
}
