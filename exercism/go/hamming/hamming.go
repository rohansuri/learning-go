package hamming

import "errors"

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("hamming distance is only calculable for strands of equal length")
	}
	brune := []rune(b)
	distance := 0
	for i, arune := range a {
		if arune != brune[i] {
			distance++
		}
	}

	return distance, nil
}
