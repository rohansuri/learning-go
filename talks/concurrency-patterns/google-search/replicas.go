package google_search

import (
	"fmt"
	"time"
)

// version 3.0
// remove tail latencies by sending queries to multiple replicas
// even if some are having troubles, one good one will come back in time

func ConcurrentGoogleQueryToReplicas(query string) []string {
	var results []string

	ch := make(chan string)

	// basic fan in, combining results into one channel
	go func() { ch <- First(query, Web, Web, Web) }()
	go func() { ch <- First(query, Image, Image, Image) }()
	go func() { ch <- First(query, Video, Video, Video) }()

	timeout := time.After(50 * time.Millisecond)

	for i := 0; i < 3; i++ {
		select {
		case result := <-ch:
			results = append(results, result)
		case <-timeout:
			fmt.Println("timed out")
			break
		}

	}

	return results
}

func First(query string, replicas ...Search) string {
	return <-fanIn(query, replicas...)
}
