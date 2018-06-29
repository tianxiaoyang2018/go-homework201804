package main

import "fmt"

func main() {
	arr := [4]int{1,2,3,4}
	fmt.Println(arr,len(arr), cap(arr))

	slice := make([]int, 1)
	fmt.Println(slice, len(slice), cap(slice))
	slice = append(slice, 1)
	fmt.Println(slice, len(slice), cap(slice))
	slice = append(slice,2)
	fmt.Println(slice, len(slice), cap(slice))
	slice = append(slice,3)
	fmt.Println(slice, len(slice), cap(slice))
	slice = append(slice,4,5,6,7,8)
	fmt.Println(slice, len(slice), cap(slice))
	
	
}
