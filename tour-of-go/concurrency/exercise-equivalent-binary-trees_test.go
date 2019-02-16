package concurrency

import (
	"testing"

	"golang.org/x/tour/tree"
)

func TestWalk(t *testing.T) {
	ch := make(chan int)
	tree := tree.New(1) // 1, 2, 3, ... 10

	go Walk(tree, ch)

	for i := 1; i <= 10; i++ {
		if got := <-ch; i != got {
			t.Errorf("Expected: %v, Got: %v", i, got)
		}
	}
}

func TestSame(t *testing.T) {
	table := []struct {
		t1       *tree.Tree
		t2       *tree.Tree
		expected bool
	}{{
		tree.New(1),
		tree.New(1),
		true,
	}, {
		tree.New(1),
		tree.New(2),
		false,
	}}

	for _, test := range table {
		if Same(test.t1, test.t2) != test.expected {
			t.Fail()
		}
	}
}
