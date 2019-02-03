package stringutil

func Reverse(s string) string {
	// runes are an alias for int32 to cover entire Unicode space
	// see https://golang.org/ref/spec#Conversions
	chars := []rune(s)
	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
		// TODO: why does this work?
		chars[i], chars[j] = chars[j], chars[i]
	}
	return string(chars)
}
