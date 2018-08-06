package main

import "fmt"

type Node struct {
	value string
	next  *Node
}

func CreateList(arr []string) *Node {
	if len(arr) == 0 {
		return nil
	}
	var head *Node = nil
	for i := len(arr) - 1; i >= 0; i-- {
		temp := &Node{value: arr[i], next: head}
		//fmt.Println("tmep=",temp)
		head = temp
	}
	//fmt.Println(head)
	return head
}
func main() {
	list := CreateList([]string{"a", "b", "c", "d"})
	//fmt.Println(list)
	for ; list != nil; list = list.next {
		fmt.Println(list.value)
	}
}
