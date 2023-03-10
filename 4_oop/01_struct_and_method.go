package tree

import "fmt"

/**
1. Golang does not support: inheritance & polymorphism
2. so there are no class but only struct
3.
	type <structName> struct {
		<param names> <types>
		...
	}
4. see different ways of create struct
5. no matter pointer or object, just use . to access
6. check createNode function, in golang, address of local variable can also be returned
7. for Golang, garbage collection + compiler will take care of all stuff, no matter it is on stack or heap
	- if some variable only used within functions, compiler might try to put that on stack
	- if it has used a bit, compiler will try to put that on stack
	- compiler will decide where to put it
	- even stuff created on heap, it doesn't need tobe GC manually
8. For struct functions (check print()):
func (<param_receiver> structType) <func_name>(<params...>) {
	... to define functions
}
then:
<structObject>.<func_name>() to call the function
9. need to define pointer/value as receiver: func (<param_receiver> *structType)... to allow changes happen to struct, otherwise, it will cause copy
10. nil pointer can also call methods!!! But access/set values from/to it will cause errors
*/

type TreeNode struct {
	value       int
	left, right *TreeNode
}

func (node TreeNode) print() {
	fmt.Println(node.value)
}

func (node *TreeNode) setValuePtr(value int) {
	if node == nil {
		fmt.Println("Setting value to nil, ignored.")
	} else {
		node.value = value
	}
}

func (node TreeNode) setValueOri(value int) {
	node.value = value
}

func createNode(value int) *TreeNode {
	return &TreeNode{value: value}
}

func (node *TreeNode) traverse() {
	if node == nil {
		return
	}
	node.left.traverse()
	node.print()
	node.right.traverse()
}

func main() {
	// create
	var root TreeNode                   // all values zero value
	root = TreeNode{value: 3}           // partial value
	root.left = &TreeNode{}             // all values zero value
	root.right = &TreeNode{5, nil, nil} // create with initial value
	root.right.left = new(TreeNode)     // new will return *Type
	root.left.right = createNode(2)     // function address here (usually return an address)
	root.print()

	// pass by pointer can change the value
	root.left.right.setValuePtr(4)
	root.left.right.print()

	root.left.right.setValueOri(5)
	root.left.right.print()

	nodes := []TreeNode{ // can create as other variables
		{value: 3},
		{},
		{6, nil, &root},
	}

	var emptyNode *TreeNode
	emptyNode.setValuePtr(1)
	emptyNode = &root
	emptyNode.setValuePtr(1)

	fmt.Println(nodes)

	/**
		1
	   / \
	  6   3
	 / \ / \
	5  4 2  7
	*/
	r := TreeNode{value: 1,
		left:  &TreeNode{value: 6, left: &TreeNode{value: 5}, right: &TreeNode{value: 4}},
		right: &TreeNode{value: 3, left: &TreeNode{value: 2}, right: &TreeNode{value: 7}}}

	r.traverse()

}
