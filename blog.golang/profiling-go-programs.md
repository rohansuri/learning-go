Notes from Go's official [blog post](https://blog.golang.org/profiling-go-programs) on Profiling Go programs.

#### Intro to the package
* `runtime/pprof` package contains functions to explicitly enable CPU and memory profile.

* They take a file to dump the profile data.

* Idiomatic usage is to enable/disable profiling by defining own flags.

* If using the benchmark methods then out of the box flags are already available:  
  -cpuprofile, -memprofile

* The profile data is examined by `go tool pprof process_binary dump.prof` command.
  
  pprof tool takes you into command mode.

  listing a few commands:  
  `topN` - shows top N samplings CPU or memory  
  `web` - generates a graph of code paths, time spent in each call  
  `weblist` - gives annotated source code with assembly to further drill down into asm instructions  
  `list <function-name>` - annotated source code with profiling data

* `net/http/pprof` creates HTTP handlers that can be invoked to create and download profiles at runtime.

  /debug/pprof/profile  
  /debug/pprof/heap

* Go program stops about 100 times in a second to record a CPU sample.


#### Investigating compute time

* Run `topN` to get the list of functions where most of the time is being spent.

    Columns:  
    \- # of samples in which function is running  
    \- % for the same  
    \- cumulative total of the listing based on column 2  
    \- # of samples in which function is running or present in codepath i.e. one of the functions it called is running  
    \- % for the same  

    By default sorted by column 1, 2 which are the important columns to look at.

    ```
    (pprof) top10
    Total: 2525 samples
        298  11.8%  11.8%      345  13.7% runtime.mapaccess1_fast64
        268  10.6%  22.4%     2124  84.1% main.FindLoops
    ```
    Shows that 11.8% of the samples collected had mapaccess being run.

* Pick a candidate function to investigate.

* Run `web <function-name>` which creates a graph of codepaths and calls leading upto the target function.

  In the blog the function to investigate was `runtime.mapaccess1_fast64`

* Since the hotspot function is a library function, using the graph trace the user code calling into it.

* List the source for that function using `list <function-name>`

  Shows out of total X samples collected, how many samples had a particular source code line running and was present in the sample's codepaths.

  ```
  1     37  242:     number[currentNode] = current
  ... lines skipped...
  9    152  246:             if number[target] == unvisited {
  ... lines skipped...
  7     59  250:     last[number[currentNode]] = lastid
  ```
* Note which lines in particular appear in most of the samples, meaning where most time is spent.

* Above example shows a lot of samples spending time in map accesses.

  Therefore we need to optimise our usage of map.
  Either do less frequent lookups or use much efficient data structure.
  Here a set backed by slice was enough and improved performance.

* Run topN again and pick another function to inspect.

* `runtime.mallocgc` accounts for most of the run time.
   It might mean the program is allocating too much memory.
   
   Resolutions would be to pool/reuse the same objects and avoiding unnecessary allocations/garbage production.

#### Investigating memory

* To understand which call sites are making the most allocations we generate a memory profile using `memprofile`.

  It generates a .mprof data file.

* Aim is to find the user code which is making allocations.

* Use `topN` and pick a candidate function for investigation.

* `list <function-name>` to list it's source and respective allocations done.

* `list --inuse_objects` gives count of allocations rather than size in bytes.

* `weblist mallogc ` to generate a graph of allocations.

* `weblist --nodefraction=0.1 mallogc` to only include code paths present in more than 0.1 fraction of samples.

* Optimise usage of your objects by reusing them.

#### Conclusion

In this case, the same bottleneck of our map usage
was accounting for both an increase in cpu time as well as lot of memory usage.


#### Questions

* Do we have anything like an external profiler?  
  To turn the knob on at sometime later at runtime maybe?



