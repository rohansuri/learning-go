package google_search

import (
	"fmt"
	"time"
)

/*

no pre allocation of results as make([]string, 3)
since we could get lesser than 3 results
therefore it is best to manipulate the slice using append that appropriately
adjusts the length and capacity.
this helps the caller to iterate only upto length elements (actual appended elements).
else with pre allocation, length = cap = 3 at all times even if we got no results!

*/

// version 2.1
func ConcurrentGoogleWithTimeout(query string) []string {
	var results []string

	ch := fanIn(query, Web, Image, Video)

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
