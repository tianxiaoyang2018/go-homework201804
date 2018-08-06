package main

import "fmt"

func sum(values []int, resultChan chan int) {
	sum := 0
	for _, value := range values {
		sum += value
	}
	resultChan <- sum
}

func main() {
	fmt.Println("测试channel用法")
	values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	resulChan := make(chan int, 3)
	go sum(values[:len(values)/2], resulChan)
	go sum(values[len(values)/2:], resulChan)
	go sum(values[len(values)/3:], resulChan)
	sum1, sum2, sum3 := <-resulChan, <-resulChan, <-resulChan
	fmt.Println("result", sum1, sum2, sum3)
}
