package crt

import (
	"fmt"
	"strings"

	colour "github.com/fatih/color"
	lang "github.com/mt1976/crt/language"
)

// The upcase function in Go converts a string to uppercase.
// upcase converts a string to uppercase.
func upcase(s string) string {
	return strings.ToUpper(s)
}

func downcase(s string) string {
	return strings.ToLower(s)
}

// The function `trimRepeatingCharacters` takes a string `s` and a character `c` as input, and returns
// a new string with all consecutive occurrences of `c` trimmed down to a single occurrence.
func trimRepeatingCharacters(s string, c string) string {

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

// italic returns a string with the italic style applied to it.
func italic(s string) string {
	return colour.New(colour.Italic).Sprintf("%v", s)
}

// boldInt returns a bolded string representation of an integer.
func boldInt(s int) string {
	return bold(fmt.Sprintf("%d", s))
}

// sQuote returns a string with the single quote symbol around it.
func sQuote(s string) string {
	return lang.SymSingleQuote + s + lang.SymSingleQuote
}

// pQuote returns a string with the square bracket symbol around it.
func pQuote(s string) string {
	return lang.SymSquareQuoteOpen + s + lang.SymSquareQuoteClose
}

// dQuote returns a string with the double quote symbol around it.
func dQuote(s string) string {
	return lang.SymDoubleQuote + s + lang.SymDoubleQuote
}

// qQuote returns a string with the square quote symbol around it.
func qQuote(s string) string {
	return lang.SymSquareQuoteOpen + s + lang.SymSquareQuoteClose
}

// isActionIn determines if the input string contains any of the specified actions.
// It is case-insensitive.
//
// Parameters:
// in: The input string to search.
// check: The list of actions to check for.
//
// Returns:
// A boolean indicating whether the input string contains any of the specified actions.
func isActionIn(in string, check ...string) bool {
	for i := 0; i < len(check); i++ {
		if strings.Contains(upcase(in), check[i]) {
			return true
		}
	}
	return false
}
