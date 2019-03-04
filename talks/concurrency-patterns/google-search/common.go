package google_search

import (
	"fmt"
	"math/rand"
	"time"
)

// Question: when to create a function type?
// when all it's methods are only dependent on the input parameters?

type Search func(query string) string

var (
	Web   = createSearch("Web")
	Image = createSearch("Image")
	Video = createSearch("Video")
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func createSearch(kind string) Search {
	return func(query string) string {
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(100)))
		return fmt.Sprintf("%s result for \"%s\"", kind, query)
	}
}
