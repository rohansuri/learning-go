package rot13reader

import (
	"fmt"
	"io"
	"unicode"
)

type rot13Reader struct {
	r io.Reader
}

func (rot13 *rot13Reader) Read(arr []byte) (int, error) {
	// fill the input byte slice by rot13 cipher of the
	// underlying reader we have
	// we return EOF when, underlying reader gives us EOF
	// and we stop until underlying reader gives EOF
	// or till we reach byte slice's length

	buffer := make([]byte, 1024)
	bytesWritten := 0
	for {
		bytesRead, err := rot13.r.Read(buffer)
		if err != nil {
			return bytesWritten, err
		}

		rotate13(buffer[:bytesRead])

		// write the buffer to the given byte slice
		bytesCopied := copy(arr, buffer[:bytesRead])
		bytesWritten += bytesCopied
		if bytesRead != bytesCopied {
			// arr is not large enough, we can't copy any more
			// let's return
			return bytesWritten, nil
		}
	}
}

func rotate13(rot13 []byte) {
	for i := 0; i < len(rot13); i++ {
		r := rune(rot13[i])
		if unicode.IsLetter(r) {

			if unicode.IsUpper(r) {
				if rot13[i]+13 > byte('Z') {
					rot13[i] = byte('A') + (rot13[i] + 13) - byte('Z') - 1
				} else {
					rot13[i] += 13
				}
			} else {
				if rot13[i]+13 > byte('z') {
					rot13[i] = byte('a') + (rot13[i] + 13) - byte('z') - 1
				} else {
					rot13[i] += 13
				}
			}
		} else {
			fmt.Printf("byte value %v not a letter, skipping rot13 for it\n", rune(rot13[i]))
		}
	}
}
