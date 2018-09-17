package main

import (
	"time"
	"fmt"
)

func main() {
	var timestamp int64 = 123
	var time time.Time = time.Unix(timestamp, 0)
	t := &time
	fmt.Println(*t)
}
func Test() {
	fmt.Println("123")
}
