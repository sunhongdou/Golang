package main

import (
	"fmt"
)

type Tree struct {
	value       int
	left, right *Tree
}

func main() {
	tree1 := Tree{
		value: 1,
		left:  nil,
		right: nil,
	}
	tree2 := Tree{
		value: 2,
		left:  nil,
		right: nil,
	}
	fmt.Println(tree1)
	//结构体比较
	fmt.Println(tree1 == tree2)
}
