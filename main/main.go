package main

import (
	"errors"
	"fmt"
	"strings"
)

func test() (string, error) {
	return "abc", errors.New("123")
}
var err1 = errors.New("123 ")
func main() {
	var err error
	str, err := test()
	fmt.Println(str, err)
	fmt.Println(strings.TrimSpace(err1.Error()))
	if err.Error() == "123" {
		fmt.Println("成功")
	}
	fmt.Println(strings.Contains("Conflict ","Conflict"))
}
