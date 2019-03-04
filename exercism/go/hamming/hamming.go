package hamming

import "errors"

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("Hamming distance is only calculable for strands of equal length.")
	}

	distance := 0

	other := []rune(b)

	// range gives us runes (int32)
	// so both strings need to be runes
	for i, str := range a {
		if str != other[i] {
			distance++
		}
	}

	return distance, nil
}
