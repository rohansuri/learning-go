package concurrency_patterns

import (
	"fmt"
	"time"
)

// On every select we block over the two communication operations
// one being communication by messi, other being a timeout communication by sort of a timer
func timedOutCommunication(){
	ch := boring("Messi")

	for {
		select {
		case item := <-ch:
			fmt.Println(item)
		case <-time.After(time.Second):
			fmt.Println("You're too slow. Quitting.")
			return
		}
	}

}
