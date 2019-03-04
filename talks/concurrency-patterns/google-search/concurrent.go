package google_search

// version 2.0
func ConcurrentGoogleUsingChannels(query string) []string {
	results := make([]string, 3)

	ch := fanIn(query, Web, Image, Video)

	for i := 0; i < 3; i++ {
		results[i] = <-ch
	}

	return results
}

func fanIn(query string, searches ...Search) <-chan string {
	ch := make(chan string)

	for _, search := range searches {
		go func(search Search) { ch <- search(query) }(search)
	}

	return ch
}
