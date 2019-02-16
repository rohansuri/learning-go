package concurrency

import (
	"fmt"
)

// unrelated to tour-of-go, just for practise

type completeBinaryTree struct {
	left  *completeBinaryTree
	right *completeBinaryTree
	value int
}

/*

1, 2, 3, 4, 5

		1


		1
	   /
	  2

	    1
	   / \
	  2   3

		1
	   / \
	  2   3
	 /
	4

so on...

level wise fill up all the nodes

*/

func newTree(arr []int) *completeBinaryTree {

	if len(arr) < 1 {
		return &completeBinaryTree{}
	}

	root := &completeBinaryTree{value: arr[0]}

	queue := make([]*completeBinaryTree, 0)

	queue = append(queue, root)

	for i := 1; len(queue) != 0; {
		// get from queue
		// pick next two elements from array
		// make them my children
		// and add them to queue

		node := queue[0]  // peek
		queue = queue[1:] // pop

		fmt.Println("Popped ", node.value)

		if i < len(arr) {
			node.left = &completeBinaryTree{value: arr[i]}
			queue = append(queue, node.left)
			i++
		}

		if i < len(arr) {
			node.right = &completeBinaryTree{value: arr[i]}
			queue = append(queue, node.right)
			i++
		}
	}

	return root
}

// level order traversal (has to be for a complete binary tree)
func (t *completeBinaryTree) traverse() []int {
	if t == nil {
		return []int{}
	}

	result := make([]int, 0)

	queue := make([]*completeBinaryTree, 0)

	queue = append(queue, t)

	for len(queue) != 0 {
		node := queue[0]  // peek
		queue = queue[1:] // pop

		result = append(result, node.value)

		if node.left != nil {
			queue = append(queue, node.left)
		} else {
			continue // since it is a complete binary free, which is filled left first on each level
		}

		if node.right != nil {
			queue = append(queue, node.right)
		}

	}

	return result
}
