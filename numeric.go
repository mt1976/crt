package crt

import (
	"math/rand"
	"strconv"
	"unicode"
)

// The isInt function checks if a given string consists only of digits.
// isInt checks if a given string consists only of digits.
func isInt(in string) bool {
	if in == "" {
		return false
	}
	for _, c := range in {
		if !unicode.IsDigit(c) {
			return false
		}
	}
	return true
}

func coinToss() bool {
	// This function simulates a coin toss.
	var coinSides = 2
	return rand.Intn(coinSides) != 0
}

func toInt(in string) int {
	// This function converts a string to an integer.
	i, err := strconv.Atoi(in)
	if err != nil {
		return 0
	}
	return i
}

func toString(in int) string {
	// This function converts an integer to a string.
	return strconv.Itoa(in)
}
