package main

import (
	"golang.org/x/tour/reader"
)

type MyReader struct{}

// Exercise: Add a Read([]byte) (int, error) method to MyReader.

// we never return io.EOF, since we're asked to return an infinite stream
func (reader *MyReader) Read(arr []byte) (int, error) {
	//fmt.Printf("length of byte array given: %v\n", len(arr))
	for i := 0; i < len(arr); i++ {
		arr[i] = byte('A')
	}
	return len(arr), nil
}

func main() {
	reader.Validate(&MyReader{})
}
