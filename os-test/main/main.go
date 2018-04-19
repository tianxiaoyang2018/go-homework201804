package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	fileInfo, err := os.Stat("/Users/tianxiaoyang/Desktop/go-test/hello.go")
	fmt.Println(fileInfo.Name(), fileInfo.Size())
	fmt.Println(err)

	all, err := ioutil.ReadFile("/Users/tianxiaoyang/Desktop/go-test/hello.go")
	fmt.Println(string(all))
	fmt.Println(err)

	host, err := os.Hostname()
	fmt.Println("host=", host)
}
