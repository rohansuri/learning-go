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

	go func(){
		for i := 1; ; i++ {
			ch <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1000)))
		}
	}()

	return ch
}
