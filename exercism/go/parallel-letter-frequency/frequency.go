package letter

import (
	"runtime"
	"strings"
)

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s *string) FreqMap {
	m := FreqMap{}
	for _, r := range *s {
		m[r]++
	}
	return m
}

/*
Counting letters is a CPU bound task, therefore we should only create goroutines as many as CPU cores.
Moreover the sub task size for each routine should be large enough,
so that the cost of frequency count merge from another routine is amortized against the counting in parallel speed gain.

we could potentially even spawn multiple routines that
merge the counts.

essentially a MapReduce M * R model.
M large
R small
and a final reducer phase to merge all Rs.

how can the expense be measured?

original task T
split amongst n workers
so each gets T/n size
which means after n workers finish, there'll be a single merge of those n sub problems.

mapper cost = O(T/n)
reducer cost = no of sub solutions * time taken to aggregate back a single solution
             = n * T/n
             = O(T)

so complexity stays the same
where is the gain?
*/

// ConcurrentFrequency counts the frequency of each rune in the given texts concurrently and returns the aggregate
// as a FreqMap.
func ConcurrentFrequency(texts []string) FreqMap {
	text := strings.Join(texts, "")

	noOfTasks := runtime.NumCPU() - 1 // keep the main go routine for merging counts

	taskLength := len(text) / noOfTasks //int(math.Ceil(float64(len(text)) / float64(noOfTasks)))

	// no need to parallelize
	// also include a minimal sub problem size to pick the number of tasks?
	if taskLength < 0 {
		noOfTasks = 1
	}

	// buffered channel to let tasks finish up without waiting for main to consume it's result.
	// since our main goroutine could be busy updating the merged frequency count
	taskResults := make(chan FreqMap, noOfTasks)

	// divide
	for i := 0; i < noOfTasks-1; i++ {
		task := text[i*taskLength : (i+1)*taskLength]
		go func() {
			taskResults <- Frequency(&task)
		}()
	}

	go func() {
		lastTask := text[(noOfTasks-1)*taskLength:]
		taskResults <- Frequency(&lastTask)
	}()

	// conquer
	return mergeCounts(taskResults)
}

func mergeCounts(ch chan FreqMap) (freqMap FreqMap) {
	freqMap = <-ch // we'd have at least one task

	for i := 0; i < cap(ch)-1; i++ {
		freq := <-ch

		for letter, count := range freq {
			freqMap[letter] += count
		}
	}
	return
}
