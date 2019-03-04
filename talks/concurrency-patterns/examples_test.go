package concurrency_patterns

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
)

func init(){
	go func() {
		if err := http.ListenAndServe("localhost:6060", nil); err != nil {
			log.Fatalln("Failed to start pprof debug HTTP handler, quitting.", err)
		}
	}()
}

func ExampleGenerator(){
	ch := boring("boring")

	for i := 1; i <= 3; i++ {
		output := <-ch
		fmt.Println(output)
	}

	// Output:
	// boring 1
	// boring 2
	// boring 3
}

// GeneratorAsAService demonstrates thinking of generators as independent instances
// of a service.
// Spawning multiple instances and communicating with them using channels.
func ExampleGeneratorAsAService(){
	messi := boring("Messi")
	cristiano := boring("Cristiano")

	for i := 0; i < 3; i++ {

		fmt.Println(<-messi)
		fmt.Println(<-cristiano)
	}

	// Output:
	// Messi 1
	// Cristiano 1
	// Messi 2
	// Cristiano 2
	// Messi 3
	// Cristiano 3
}

func ExampleFanIn(){
	ch := fanIn(boring("Messi"), boring("Cristiano"))

	for i := 0; i < 6; i++ {
		fmt.Println(<-ch)
	}

	// Unordered output:
	// Messi 1
	// Cristiano 1
	// Messi 2
	// Cristiano 2
	// Messi 3
	// Cristiano 3
}