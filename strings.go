package crt

import (
	"strings"

	l "github.com/mt1976/crt/language"
)

// The Upcase function in Go converts a string to uppercase.
// Upcase converts a string to uppercase.
func Upcase(s string) string {
	return strings.ToUpper(s)
}

// The function `TrimRepeatingCharacters` takes a string `s` and a character `c` as input, and returns
// a new string with all consecutive occurrences of `c` trimmed down to a single occurrence.
func TrimRepeatingCharacters(s string, c string) string {

	result := ""
	lenS := len(s)

	for i := 0; i < lenS; i++ {
		if i == 0 {
			result = string(s[i])
		} else {
			if string(s[i]) != c || string(s[i-1]) != c {
				result = result + string(s[i])
			}
		}
	}
	return result
}

func Bold(s string) string {
	return l.TextStyleBold + s + l.TextStyleReset
}

func SQuote(s string) string {
	return l.SymSingleQuote + s + l.SymSingleQuote
}

func PQuote(s string) string {
	return l.SymOpenBracket + s + l.SymCloseBracket
}
