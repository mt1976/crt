package support

import (
	"fmt"
	"math/rand"
	"time"
)

// The function roundFloatToTwo rounds a float64 number to two decimal places.
// func roundFloatToTwo(input float64) float64 {
// 	// round float64 to 2 decimal places
// 	rtnVal := math.Round(input*100) / 100

// 	return rtnVal
// }

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

// The function DateString returns the current date in the format "dd/mm/yy".
func DateString() string {
	now := time.Now()
	return fmt.Sprintf("%v", now.Format("02/01/06"))
}

// The TimeString function returns the current time in the format "15:04:05".
func TimeString() string {
	now := time.Now()
	return fmt.Sprintf("%v", now.Format("15:04:05"))
}

// The DateTimeString function returns a string that combines the time and date strings.
func DateTimeString() string {
	return TimeString() + " " + DateString()
}

// The RandomIP function generates a random IP address in IPv4 format.
func RandomIP() string {
	// Generate a random IP address in ipv4 format
	//
	// Example: 123.456.789.012
	//
	// Returns:
	// 	string: A random IP address
	//
	// Usage:
	// 	ip := randomIP()
	// 	fmt.Println(ip)
	//
	ip1 := RandomNumber(1, 255)
	ip2 := RandomNumber(1, 255)
	ip3 := RandomNumber(1, 255)
	ip4 := RandomNumber(1, 255)

	return fmt.Sprintf("%v.%v.%v.%v", ip1, ip2, ip3, ip4)
}

// The RandomMAC function generates a random MAC address in the format of 00:00:00:00:00:00.
func RandomMAC() string {
	// Generate a random MAC address in the format of 00:00:00:00:00:00
	//
	// Returns:
	// 	string: A random MAC address
	//
	// Usage:
	// 	mac := randomMAC()
	// 	fmt.Println(mac)
	//
	mac1 := fmt.Sprintf("%02x", RandomNumber(0, 255))
	mac2 := fmt.Sprintf("%02x", RandomNumber(0, 255))
	mac3 := fmt.Sprintf("%02x", RandomNumber(0, 255))
	mac4 := fmt.Sprintf("%02x", RandomNumber(0, 255))
	mac5 := fmt.Sprintf("%02x", RandomNumber(0, 255))
	mac6 := fmt.Sprintf("%02x", RandomNumber(0, 255))

	return fmt.Sprintf("%v:%v:%v:%v:%v:%v", mac1, mac2, mac3, mac4, mac5, mac6)
}

// The RandomPort function generates a random port number between 1 and 65535.
func RandomPort() int {
	// Generate a random port number between 1 and 65535
	//
	// Returns:
	// 	int: A random port number
	//
	// Usage:
	// 	port := randomPort()
	// 	fmt.Println(port)
	//
	return RandomNumber(1, 65535)
}

// The RandomNumber function generates a random number within a given range.
func RandomNumber(min int, max int) int {
	// Generate a random number between the given range
	//
	xx := rand.Intn(max-min+1) + min

	return xx
}
