// Package twofer is short for two for one i.e.
// when two articles are being sold for the price of one.
package twofer

import "fmt"

// ShareWith returns a string expressing that the twofer
// will be shared by the caller.
func ShareWith(name string) string {
	if len(name) == 0 {
		name = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}
