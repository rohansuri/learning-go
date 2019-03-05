package google_search

import (
	"fmt"
	"sort"
	"testing"
)

var query = "Messi"

var expected = []string{"Video result for \"" + query + "\"",
	"Web result for \"" + query + "\"",
	"Image result for \"" + query + "\""}

// empty struct takes 0 bytes
var expectedMap = make(map[string]struct{})

func init() {
	for _, element := range expected {
		expectedMap[element] = struct{}{}
	}
}

func BenchmarkSynchronousGoogle(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SynchronousGoogle(query)
	}
}

func BenchmarkConcurrentGoogleUsingClosures(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ConcurrentGoogleUsingClosures(query)
	}
}

func BenchmarkConcurrentGoogleUsingChannels(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ConcurrentGoogleUsingChannels(query)
	}
}

type Google func(query string) []string

func TestVersionsOfGoogle(t *testing.T) {
	google := []Google{ConcurrentGoogleUsingChannels, ConcurrentGoogleUsingClosures, SynchronousGoogle}

	for i, g := range google {
		results := g(query)
		assertEqualSlices(results, expected, t)
		if t.Failed() {
			t.Errorf("Test case with function at index %d failed", i)
		}
	}
}

func assertEqualSlices(got, expected []string, t *testing.T) {
	sort.Strings(expected)
	sort.Strings(got)

	if len(got) != len(expected) {
		t.Errorf("Result size expected %d, got %d\n", len(expected), len(got))
	}

	for i, result := range got {
		if expected[i] != result {
			t.Errorf("Expected result: %s, Got: %s", expected[i], result)
		}
	}
}

func TestConcurrentGoogleWithTimeout(t *testing.T) {
	results := ConcurrentGoogleWithTimeout(query)
	fmt.Println(results)
	// every result must be one of the expected results
	for _, result := range results {
		if _, isPresent := expectedMap[result]; !isPresent {
			t.Fatalf("Unexpected result %s", result)
		}
	}
}

func TestConcurrentGoogleWithReplicas(t *testing.T) {
	results := ConcurrentGoogleQueryToReplicas(query)
	fmt.Println(results)
	// every result must be one of the expected results
	for _, result := range results {
		if _, isPresent := expectedMap[result]; !isPresent {
			t.Fatalf("Unexpected result %s", result)
		}
	}
}
