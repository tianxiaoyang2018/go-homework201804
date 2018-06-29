package main

import (
	"regexp"
	"fmt"
)

func main() {
	a := "\\/"
	fmt.Println(a)
	str := "啊啊    啊啊  和"
	pat := "[ ]+"
	re,_ := regexp.Compile(pat)

	if ok, _ := regexp.Match(pat, []byte(str)); ok {
		fmt.Println("match found")
	}

	str = re.ReplaceAllString(str, " ")
	fmt.Println(str)
}
