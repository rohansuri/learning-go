package concurrency_patterns


// fanIn is a better way to fan in/multiplex over multiple channels for input
// using select
func fanIn(ch1, ch2 <- chan string) <- chan string {
	ch := make(chan string)

	go func() {
		for {
			select {
			case item := <-ch1:
				ch <- item
			case item := <-ch2:
				ch <- item
			}
		}
	}()

	return ch
}
