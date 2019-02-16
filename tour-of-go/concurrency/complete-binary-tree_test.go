package concurrency

import (
	"testing"
)

func TestTree(t *testing.T) {
	// from an array create a tree

	table := [][]int{
		{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		{1},
	}

	for _, arr := range table {
		tree := newTree(arr)
		result := tree.traverse()

		for i, value := range arr {
			if value != result[i] {
				t.Errorf("At index %v, Expected: %v, Got: %v", i, arr[i], result[i])
			}
		}
	}

}

func TestNil(t *testing.T) {
	var tree *completeBinaryTree

	if result := tree.traverse(); result == nil || len(result) != 0 {
		t.Errorf("Expected %v to be not nil and it's len %v to be 0", result, len(result))
	}
}
