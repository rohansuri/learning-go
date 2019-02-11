package slices

import (
	"fmt"
	"testing"

	"golang.org/x/tour/pic"
)

func TestPic(t *testing.T) {
	pic.Show(Pic)
}

func inc(arr []int) {
	for i := 0; i < len(arr); i++ {
		arr[i]++ // works, since we're modifying the underlying array values pointed to by the slice
	}
}

func myAppend(arr []int) {
	arr = append(arr, 2)                  // we're out of capacity, resize, new slice returned
	fmt.Println("Inside myAppend: ", arr) // [1 1 1 1 1 1 1 1 1 1 2]
}

func TestSlicePassByValue(t *testing.T) {

	arr := make([]int, 10)                   // len = cap = 10
	fmt.Println("Before calling inc: ", arr) // [0 0 0..] 10 times
	inc(arr)
	fmt.Println("After calling inc: ", arr) // [1 1 1 ..] 10 times

	fmt.Println("Before appending :", arr)
	myAppend(arr)
	fmt.Println("After appending :", arr) // [1 1 1 1 1 1 1 1 1 1]

	// we don't see the 2 that got appended, since that's a new slice that was returned
	// after allocating a new array with a larger size.
	// but in this caller, we still hold the old slice.
}
