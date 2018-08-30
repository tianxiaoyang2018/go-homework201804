package service

import "fmt"

func init() {
	fmt.Println("service实例化")
}

func Get() int {
	return 1
}