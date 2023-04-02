package main

import "fmt"

type treeNode struct {
	Val   int
	Left  *treeNode
	Right *treeNode
}

func (node *treeNode) TraverseFunc(f func(node *treeNode)) {
	f(node)
	if node.Left != nil {
		node.Left.TraverseFunc(f)
	}
	if node.Right != nil {
		node.Right.TraverseFunc(f)
	}
}

func (node *treeNode) TraverseWithChannel() chan *treeNode {
	out := make(chan *treeNode)
	go func() {
		node.TraverseFunc(func(node *treeNode) {
			out <- node
		})
		close(out)
	}()
	return out
}

func main() {
	r := treeNode{Val: 1,
		Left:  &treeNode{Val: 6, Left: &treeNode{Val: 5}, Right: &treeNode{Val: 4}},
		Right: &treeNode{Val: 3, Left: &treeNode{Val: 2}, Right: &treeNode{Val: 7}}}

	c := r.TraverseWithChannel() // pass all node into channel
	for n := range c {
		fmt.Println(n.Val)
	}
}

/*
1. Here pass all node into the channel
2. Since we are using n := range c, it will not exit until find close statement, this ensures only when goroutine within
TraverseWithChannel finished for will finish.
*/
