package support

import "unicode"

// The IsInt function checks if a given string consists only of digits.
// IsInt checks if a given string consists only of digits.
func IsInt(s string) bool {
	for _, c := range s {
		if !unicode.IsDigit(c) {
			return false
		}
	}
	return true
}
