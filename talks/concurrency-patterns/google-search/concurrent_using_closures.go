package google_search

import "sync"

// not done in the talk, but my attempt to look and analyze performance
// the quick 1 second benchmark says this is slower than using channels
func ConcurrentGoogleUsingClosures(query string) []string {
	results := make([]string, 3)

	wg := sync.WaitGroup{}
	wg.Add(3)

	// is doing this thread safe?
	// assigning slice indexes from a go routine running closure.
	// should be since all go routines are using different index of the slice
	// and the reader of the entire slice -- the main go routine
	// if wg.Wait() gives happens before then this should be thread safe
	go func(){
		results[0] = Web(query)
		wg.Done()
	}()

	go func(){
		results[1] = Image(query)
		wg.Done()
	}()

	go func(){
		results[2] = Video(query)
		wg.Done()
	}()

	wg.Wait()

	return results
}
