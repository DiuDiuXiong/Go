package mainn

/**
In order to extend existing struct/ primitive data types, inheritance not allowed in golang.
So in Golang there are two methods:
- Define in another name (don't have to be a struct, rename it in another way and write functions above it)
- Composition (let new type have one variable/pointer to the old type)
- Embedding (see example, type a/many type without parameter namespace). It will:
	- get public methods/variables those pointers point to (access directly via .)
	- allow embed more than one
	- after embedding, can override, after override:
		<final_struct>.<embedded_struct>.<override_function_name>(): call original function
		<final_struct>.<override_function_name>(): call override function, actually its just shadowing
	- note that through embedding, the child struct cannot be upper cast to father struct, it is just a simplified version of expression/grammar
	- to achieve upper cast, Golang can achieve that through 接口
*/

import (
	"Go/4_oop"
	"fmt"
	"math"
)

// By composition
type MyTreeNode struct {
	node *tree.TreeNode
}

func (myNode *MyTreeNode) PrintSomethingElse() {
	if myNode == nil {
		return
	}

	//myNode.node.CanDefineHere()
	fmt.Println("Can Print Something Else")
}

// By re-naming

type Queue []int

func (q *Queue) Push(v int) {
	*q = append(*q, v)
}

func (q *Queue) Pop() int {
	if len(*q) != 0 {
		head := (*q)[0]
		(*q) = (*q)[1:]
		return head
	} else {
		return math.MinInt
	}
}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}

// by embedding
type EmbeddingTreeNode struct {
	*tree.TreeNode // Embedding
	*MyTreeNode
}

func call_embedding_tree_node_function(n EmbeddingTreeNode) {
	n.TreeNode.CanDefineHere()        // call function by calling type directory
	n.CanDefineHere()                 // Or by calling public stuff from original struct directory
	n.PrintSomethingElse()            // can embed more than one
	n.MyTreeNode.PrintSomethingElse() // can embed more than one
}
