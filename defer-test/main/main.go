package main

import "fmt"

func test1() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}

func main() {

}
