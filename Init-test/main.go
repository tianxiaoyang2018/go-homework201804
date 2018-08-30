package main

import (
	"github.com/tianxiaoyang2018/go-test/Init-test/service"
	"fmt"
)

func main() {
	num := service.Get()
	fmt.Println(num)
}