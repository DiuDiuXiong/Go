package main

import (
	tree "Go/4_oop"
	mainn "Go/4_oop/03_extend_struct_primitive_type"
	"fmt"
)

/*
import (
	tree "Go/4_oop"
	"fmt"
)*/

func main() {
	r := tree.TreeNode{}
	r.CanDefineHere()
	fmt.Println("Main here.")

	rr := mainn.MyTreeNode{}
	rr.PrintSomethingElse()

	q := mainn.Queue{1}
	q.Push(1)
	q.Push(2)
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())
}

