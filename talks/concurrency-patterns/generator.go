package concurrency_patterns

import (
	"fmt"
	"math/rand"
	"time"
)

/*

Generator pattern:

a function starting off a go routine and generating values in it,
while returning a channel back to caller for communication.
 */

// Boring arranges a channel for the caller to receive natural numbers.
func boring(msg string) <- chan string {
	ch := make(chan string)

	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	go func(){
		for i := 1; ; i++ {
			ch <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(r.Intn(2 * 1e3)) * time.Millisecond)
		}
	}()

	return ch
}
