package concurrency_patterns

/*

every gopher is hand in hand forming a chain.
it receives an item from prev gopher and sends it to next gopher
that means it needs a communication handle for prev and next gopher

should that handle be a separate channel for each gopher?

a -> b -> c -> d

b should receive from a and send to c
so everyone could have a single channel where they expect to receive something
and then put the same item to their next gopher's channel

so every gopher should have it's own receiver channel and a handle to next gopher's receiver channel

 */

 func daisyChain(gophers int) int {
	prev := make(chan int)

	first := prev

	for i := 0; i < gophers; i++ {
		next := make(chan int)
		// what each gopher will do
		go func(prev, next chan int) {
			next <- 1 + <- prev
		}(prev, next)

		prev = next
	}

	first <- 0
	return <-prev
}
