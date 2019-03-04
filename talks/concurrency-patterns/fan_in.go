package concurrency_patterns

/*

fan in pattern:
multiplex and combine results from multiple channels into a single generator channel.
this enables the caller to not wait or process the input channels in any particular order.
whichever channel gets an item first, the caller consumes it and therefore the order is not defined
in the source code, but rather at runtime by the speed of the producers.
 */

func fanIn(ch1, ch2 <- chan string) <- chan string {
	ch := make(chan string)

	// loop here is important to continue multiplexing
	go func() { for { ch <- <-ch1 } }()
	go func() { for { ch <- <-ch2 }}()

	return ch
}
