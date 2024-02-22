package support

import (
	"math/rand"
	"strconv"
	"unicode"
)

// const RowLength int = 80
// const TitleLength int = 40
//const MenuItemLength int = 50

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

func ToInt(s string) int {
	// This function converts a string to an integer.
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return i
}

func ToString(i int) string {
	// This function converts an integer to a string.
	return strconv.Itoa(i)
}
