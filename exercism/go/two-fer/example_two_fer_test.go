package twofer

import "fmt"

// ExampleShareWith demonstrates usage of ShareWith
// with a sample input, output.
func ExampleShareWith() {
	h := ShareWith("")
	fmt.Println(h)
	// Output: One for you, one for me.
}
