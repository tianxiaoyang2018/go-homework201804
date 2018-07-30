package main

import "fmt"

type TreeNode struct {
	value int
	left *TreeNode
	right *TreeNode
}

func CreateBinayTree() *TreeNode{
	var value int
	fmt.Scanf("%d",&value)
	if value>0 {
		tmp := &TreeNode{value:value}
		tmp.left = CreateBinayTree()
		tmp.right = CreateBinayTree()
		return tmp
	}else {
		return nil
	}
}

func main() {
	tree := CreateBinayTree()
	fmt.Println(tree,*tree.left,*tree.right)
}