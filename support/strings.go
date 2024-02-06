package support

import (
	"fmt"
	"math/rand"
)

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

func RandomNumber(min int, max int) int {
	// Generate a random number between the given range
	//
	xx := rand.Intn(max-min+1) + min

	return xx
}
