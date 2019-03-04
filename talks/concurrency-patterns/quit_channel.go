package concurrency_patterns

import "fmt"

// boringQuit keeps generating till caller puts an item in the quit channel
func boringQuit(msg string, quit chan bool) <-chan string{
	ch := make(chan string)

	go func() {
		for i := 1; ; i++ {
			select {
				case <-quit: return
				case ch <- fmt.Sprintf("%s %d", msg, i):
			}
		}
	}()

	return ch
}


