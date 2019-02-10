package interfaces

import (
	"fmt"
	"testing"
)

func TestStringer(t *testing.T) {
	inputs := []struct {
		ip       IPAddr
		expected string
	}{
		{IPAddr{127, 0, 0, 1}, "127.0.0.1"},
		{IPAddr{8, 8, 8, 8}, "8.8.8.8"},
	}

	for _, input := range inputs {
		s := fmt.Sprintf("%v", input.ip)
		if s != input.expected {
			t.Errorf("Expected: %v, Got: %v", input.expected, s)
		}
	}
}
