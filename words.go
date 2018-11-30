package words

import (
	"strings"
	"unicode"
)

// Capitalize converts the first letter to uppercase in the given string
func Capitalize(s string) string {
	first := true
	return strings.Map(
		func(r rune) rune {
			if first {
				first = false
				return unicode.ToUpper(r)
			}
			return r
		},
		s)
}
