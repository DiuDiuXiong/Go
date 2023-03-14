package main

import "fmt"

type TreeNode struct {
	value       int
	left, right *TreeNode
}

func (n *TreeNode) TraverseFunc(f func(node *TreeNode)) {
	if n == nil {
		return
	}
	n.left.TraverseFunc(f)
	f(n)
	n.right.TraverseFunc(f)
}

func main() {
	r := TreeNode{value: 1,
		left:  &TreeNode{value: 6, left: &TreeNode{value: 5}, right: &TreeNode{value: 4}},
		right: &TreeNode{value: 3, left: &TreeNode{value: 2}, right: &TreeNode{value: 7}}}

	// count how many nodes are there
	nodeCount := 0
	r.TraverseFunc(func(node *TreeNode) {
		nodeCount++
	})
	fmt.Println(nodeCount)

	// calculate sum of all nodes
	sum := 0
	r.TraverseFunc(func(node *TreeNode) {
		sum += node.value
	})

	fmt.Println(sum)
}
