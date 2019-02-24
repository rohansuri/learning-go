package letter

import (
	"reflect"
	"runtime"
	"strings"
	"testing"
)

// In the separate file frequency.go, you are given a function, Frequency(),
// to sequentially count letter frequencies in a single text.
//
// Perform this exercise on parallelism using Go concurrency features.
// Make concurrent calls to Frequency and combine results to obtain the answer.

var (
	euro = `Freude schöner Götterfunken
Tochter aus Elysium,
Wir betreten feuertrunken,
Himmlische, dein Heiligtum!
Deine Zauber binden wieder
Was die Mode streng geteilt;
Alle Menschen werden Brüder,
Wo dein sanfter Flügel weilt.`

	dutch = `Wilhelmus van Nassouwe
ben ik, van Duitsen bloed,
den vaderland getrouwe
blijf ik tot in den dood.
Een Prinse van Oranje
ben ik, vrij, onverveerd,
den Koning van Hispanje
heb ik altijd geëerd.`

	us = `O say can you see by the dawn's early light,
What so proudly we hailed at the twilight's last gleaming,
Whose broad stripes and bright stars through the perilous fight,
O'er the ramparts we watched, were so gallantly streaming?
And the rockets' red glare, the bombs bursting in air,
Gave proof through the night that our flag was still there;
O say does that star-spangled banner yet wave,
O'er the land of the free and the home of the brave?`
)

func OriginalFrequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

func wrapCallToFrequency(text string, ch chan <- FreqMap){
	freq := Frequency(text)
	ch <- freq
}

func mergeCounts(ch chan FreqMap) (freqMap FreqMap){
	freqMap = make(FreqMap)

	for i := 0; i < cap(ch); i++ {
		freq := <-ch

		for letter, count := range freq {
			freqMap[letter] += count
		}
	}
	return
}

func distributeTasks(text string) chan FreqMap {
	// taskLength := int(math.Ceil(float64(len(text))/float64(runtime.NumCPU())))
	noOfTasks := runtime.NumCPU()

	taskLength := len(text) / noOfTasks

	// no need to parallelize
	if taskLength < 0 {
		noOfTasks = 1
	}

	// buffered channel to let tasks finish up without waiting for main to consume it's result.
	// since our main goroutine could be busy updating the merged frequency count
	taskResults := make(chan FreqMap, noOfTasks)

	// println("number of tasks: ", noOfTasks)
	// println("task length: ", taskLength)
	// println("size of channel:", cap(taskResults))

	for i := 0; i < noOfTasks - 1; i++ {
		task := text[i * taskLength : (i + 1) * taskLength]

		go wrapCallToFrequency(task, taskResults)
	}

	if noOfTasks > 0 {
		lastTask := text[(noOfTasks - 1) * taskLength:]
		go wrapCallToFrequency(lastTask, taskResults)
	}

	return taskResults
}

func ConcurrentFrequency(texts []string) FreqMap {
	// split total text into as many cores as you have (cpu bound task)

	text := strings.Join(texts, "")

	// println("Joined string ", text)

	taskResults := distributeTasks(text)

	return mergeCounts(taskResults)
}

func TestEdgeCases(t *testing.T){
	table := [][]string {
		{},
		{"a"},
		{strings.Repeat("a", runtime.NumCPU())},
	}

	for _, testCase := range table {

		got := ConcurrentFrequency(testCase)

		expected := OriginalFrequency(strings.Join(testCase, ""))

		if !reflect.DeepEqual(got, expected){
			t.Fatalf("Expected %v; Got %v", expected, got)
		}
	}
}

func TestConcurrentFrequency(t *testing.T) {
	seq := OriginalFrequency(euro + dutch + us)
	con := ConcurrentFrequency([]string{euro, dutch, us})
	if !reflect.DeepEqual(con, seq) {
		t.Fatal("ConcurrentFrequency wrong result")
	}
}

func TestSequentialFrequency(t *testing.T) {
	oSeq := OriginalFrequency(euro + dutch + us)
	seq := Frequency(euro + dutch + us)
	if !reflect.DeepEqual(oSeq, seq) {
		t.Fatal("Frequency wrong result")
	}
}

func BenchmarkSequentialFrequency(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Frequency(euro + dutch + us)
	}
}

func BenchmarkConcurrentFrequency(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ConcurrentFrequency([]string{euro, dutch, us})
	}
}
