package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {

	wordCount := make(map[string]int)

	words := strings.Fields(s)
	for _, word := range words {
		// if absent, then zero-value of int which is 0 is returned
		// this is good :) no need for explicit if-exists checks
		wordCount[word]++
	}

	return wordCount
}

func main() {
	wc.Test(WordCount)
}
