package main

import "fmt"

func main() {
	slice := []int{1,2}
	slice2 := append(slice,3)
	slice3 := slice2[0:2]
	slice2[0]=0
	fmt.Println(slice)
	fmt.Println(slice2)
	fmt.Println(slice3)
}
