Google I/O 2012 talk by Rob Pike on Concurrency Patterns in Go.

"Don't communicate by sharing memory, share memory by communicating"

### Concurrency

* Concurrency is the composition of independently executing computations.

* Concurrency should be thought of as a way to structure software in order to model interactions in the real world.   

  To simulate the world and the agents communicating in it with each other.

* Go's concurrency primitives make it easy to model, understand, reason about concurrency without delving into threads, locks, memory barriers, etc.  

  Rather it believes in working on a higher abstraction than these to express concurrency.

  (Just as Java abstracts object construction and it's free up when compared to C++)

* Concurrency != Parallelism  

  A concurrently modelled program may run on a single core system and may never run parallelly, but it is still concurrent.

* Based on Tony Hoare's CSP

* Don't protect shared memory by locks, mutexes, etc rather share the memory by passing it back and forth between routines.

### Goroutines

* Goroutines are independently executing functions.

* Their stacks grow and shrink in size as required.

* Many goroutines are ran multiplexed on a single kernel thread.

* It is practical to have thousands, even hundred thousands of goroutines in your program.

  Go is designed for this paradigm.

### Channels

* Primitive to communicate between goroutines.

* Other than being a send and receive operation, it is also a synchronisation operation since the communication happens with both sender and receiver in lockstep when their goroutines reach the channel operation statement and it becomes a synchronization point.

* Buffered channels remove the synchronization.

### Patterns

TODO try out all of these

* Generator (also introduces read only channel notation in function return specification)
* Multiplexing (Fan In) =->  
* Pass a channel inside a channel (to get an answer back?)
* Fan-in with select  
* Timeout each receive using `time.After`  
* Timeout whole conversation  
* Quit channel (indicate to the other goroutine to stop sending)  
* Receive on quit channel (to be sure sender has agreed to stop -- it's like we're specifying a channel protocol between the goroutines)  
* Daisy chain (Chinese whispers)  


### Google search

* Turning Google's search across Web, Image, Video from slow, sequential, failure-sensitive to fast, concurrent, replicated, robust by leveraging the concurrency patterns and primitives.

### Examples of use 

* Load balancer
* Concurrent prime sieve
* Concurrent power series


### Conclusion

* Goroutines and channels make it easy to express programs that have multiple inputs, outputs, timeouts, failures, replication, robustness, independent execution.

### Misc
* When main returns the process exits
* Designed to build systems software
* Don't overdo usage of channels, consider the right tool for the job.  
  eg. Implementing a page hits counter doesn't call for a channel.

  `sync` and `sync/atomic` packages have locks, condition variables, atomics.

* SPIN model for verification of concurrent models.


### Questions

* select vs multiple go routines  
  We can't model the situation where we want a channel value from all the cases, with multiple go routines we could do that.  
  (put an example here)

* Prioritizing cases in select statements (nested in defaults?)

* Try writing some of the patterns in Java to bring out the ease Go gives.


### Thoughts

#### Channels
Channels do offer simplicity in the way we write concurrent code.   
You don't have to think about synchronizing them or should I use a blocking queue to communicate.  
Those decisions need not bother you anymore.  
Spawning routines is cheap and programs become easier to write (take a look at the multiplexing fan in example)