package support

import (
	"math/rand"
	"unicode"
)

const MaxPageRows int = 18
const RowLength int = 80
const TitleLength int = 25
const MenuItemLength int = 50

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
