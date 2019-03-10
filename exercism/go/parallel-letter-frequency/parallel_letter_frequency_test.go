package letter

import (
	"math/rand"
	"reflect"
	"runtime"
	"strings"
	"testing"
	"time"
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

func TestEdgeCases(t *testing.T) {
	table := [][]string{
		{},
		{"a"},
		{strings.Repeat("a", runtime.NumCPU())},
	}

	for _, testCase := range table {

		got := ConcurrentFrequency(testCase)

		expected := OriginalFrequency(strings.Join(testCase, ""))

		if !reflect.DeepEqual(got, expected) {
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
	text := euro + dutch + us
	oSeq := OriginalFrequency(text)
	seq := Frequency(&text)
	if !reflect.DeepEqual(oSeq, seq) {
		t.Fatal("Frequency wrong result")
	}
}

var randRune string

func init() {
	rand.Seed(time.Now().UnixNano())
	randRune = randStringRunes(100000)
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func benchmarkSequentialFrequency(text string, b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Frequency(&text)
	}
}

func BenchmarkSequentialFrequencySmall(b *testing.B) {
	benchmarkSequentialFrequency(euro+dutch+us, b)
}

func BenchmarkSequentialFrequencyLarge(b *testing.B) {
	benchmarkSequentialFrequency(randRune, b)
}

func benchmarkConcurrentFrequency(texts []string, b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ConcurrentFrequency(texts)
	}
}

func BenchmarkConcurrentFrequencyLarge(b *testing.B) {
	benchmarkConcurrentFrequency([]string{randRune}, b)
}

func BenchmarkConcurrentFrequencySmall(b *testing.B) {
	benchmarkConcurrentFrequency([]string{euro, dutch, us}, b)
}
