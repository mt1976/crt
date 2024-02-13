package support

import (
	"math/rand"
	"unicode"
)

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

func CoinToss() bool {
	// This function simulates a coin toss.
	var coinSides = 2
	return rand.Intn(coinSides) != 0
}
