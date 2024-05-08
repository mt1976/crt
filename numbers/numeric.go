package crt

import (
	"math"
	"math/rand"
	"strconv"
	"unicode"
)

// The isInt function checks if a given string consists only of digits.
// isInt checks if a given string consists only of digits.
func IsInt(in string) bool {
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

func CoinToss() bool {
	// This function simulates a coin toss.
	var coinSides = 2
	return rand.Intn(coinSides) != 0
}

func ToInt(in string) int {
	// This function converts a string to an integer.
	i, err := strconv.Atoi(in)
	if err != nil {
		return 0
	}
	return i
}

func ToString(in int) string {
	// This function converts an integer to a string.
	return strconv.Itoa(in)
}

// The function roundFloatToTwoDPS rounds a float64 number to two decimal places.
func RoundFloatToTwoDPS(input float64) float64 {
	// round float64 to 2 decimal places
	rtnVal := math.Round(input*100) / 100

	return rtnVal
}
