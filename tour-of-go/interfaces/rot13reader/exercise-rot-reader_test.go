package rot13reader

import (
	"io/ioutil"
	"strings"
	"testing"
)

func TestRot13Reader(t *testing.T) {
	input := struct {
		ciphertext string
		plaintext  string
	}{"Lbh penpxrq gur pbqr!", "You cracked the code!"}

	s := strings.NewReader(input.ciphertext)

	r := rot13Reader{s}

	rot13, err := ioutil.ReadAll(&r)
	if err != nil {
		t.Error(err)
	}

	rot13Str := string(rot13)
	if rot13Str != input.plaintext {
		t.Errorf("Expected: %v, Got: %v", input.plaintext, rot13Str)
	}

}
