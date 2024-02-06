package support

import (
	"math"
)

// The function roundFloatToTwo rounds a float64 number to two decimal places.
func roundFloatToTwo(input float64) float64 {
	// round float64 to 2 decimal places
	rtnVal := math.Round(input*100) / 100

	return rtnVal
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
