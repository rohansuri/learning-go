package google_search


// version 1.0
func SynchronousGoogle(query string) []string {
	results := make([]string, 3)

	results[0] = Web(query)
	results[1] = Image(query)
	results[2] = Video(query)

	return results
}
