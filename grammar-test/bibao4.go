package main

import "fmt"

func adder() func(int) int {
	sum :=0
	return func(x int) int{
		sum += x
		return sum
	}
}

func add(a int) func(int) int {
	return func(b int) int {
		return a+b
	}
}

func main() {
	adder1 := adder()
	adder2 := adder()
	for i:=0;i<10;i++ {
		fmt.Println(adder1(i), adder2(-i))
	}

	add1, add2 := add(0), add(0)
	for i:=0;i<10;i++ {
		fmt.Println(add1(i),add2(-i))
	}
}
