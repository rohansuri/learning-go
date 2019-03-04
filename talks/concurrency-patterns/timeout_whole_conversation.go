package concurrency_patterns

import (
	"fmt"
	"time"
)

func timeoutWholeConversation() {
	ch := boring("Messi")

	timer := time.After(3 * time.Second)

	for {
		select {
			case item := <-ch:
				fmt.Println(item)
			case <-timer:
				fmt.Println("Times up, ending conversation")
				return
		}
	}
}
