package concurrency_patterns

import "fmt"

// boringQuitWithCleanup keeps generating till caller puts an item in the quit channel.
// Quit channel allows for explicit signal to finish generator and allows generator for cleanup.
func boringQuitWithCleanup(msg string, quit chan bool) <-chan string{
	ch := make(chan string)

	go func() {
		for i := 1; ; i++ {
			select {
			case <-quit:
				// do any cleanup, caller must be waiting for our cleanup
				// after cleanup, let's finish the program
				quit <- true
				return
			case ch <- fmt.Sprintf("%s %d", msg, i):
			}
		}
	}()

	return ch
}
