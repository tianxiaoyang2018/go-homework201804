package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		fmt.Println(2)
	}()
	fmt.Println(1)
	time.Sleep(1000)
}
